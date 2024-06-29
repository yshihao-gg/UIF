import {
  Base64
} from 'js-base64';

export function V2rayN2UIF(rawData) {
  var res = []
  try {
    res = V2rayN2UIF2(Base64.decode(rawData))
    if (res.length > 0) {
      return res
    }
  } catch (e) {
    console.log(e)
  }
  return V2rayN2UIF2(rawData);
}

// not Base64
function V2rayN2UIF2(rawData) {
  var res = []
  var protocolMap = {
    'ss:': 'shadowsocks',
    'trojan:': 'trojan',
    'trojan-go:': 'trojan',
    'hysteria2:': 'hysteria2',
    'vmess:': 'vmess'
  }
  rawData = rawData.split('\n')
  for (var item in rawData) {
    item = rawData[item]
    if (item == '') {
      continue
    }
    if (item[0] == '#') {
      continue
    }
    var hash = item.split('#')
    if (hash.length == 2) {
      if (!hash[1].includes('%')) {
        item = hash[0] + '#' + encodeURIComponent(hash[1])
      }
    }
    try {
      var url = new URL(item)
    } catch (e) {
      continue
    }
    var protocol = protocolMap[url.protocol]
    if (!['shadowsocks', 'trojan', 'vmess', 'hysteria2', 'vless'].includes(protocol)) {
      continue
    }
    if (protocol == 'shadowsocks') {
      try {
        var temp = item.split('@')
        var temp2 = temp[0].split('://')
        if (temp2[1].indexOf(':') != -1) {
          temp2 = Base64.encode(temp2[1])
          item = `ss://${temp2}@${temp[1]}`
          url = new URL(item)
        }
      } catch (e) {
        continue
      }
    }
    var transport = {
      'protocol': 'tcp',
      'tls_type': 'none',
      'tls': {},
    }
    var protocolSetting = {}
    var transportSetting = {}
    var address = ''
    var port = 443
    var tag = ''
    if (protocol == 'vmess') { //{{{
      var data = item.split('://')
      data = JSON.parse(Base64.decode(data[1]))
      url.data = data
      protocolSetting['uuid'] = data.id
      protocolSetting['alter_id'] = parseInt(data.aid)
      if ('scy' in data) {
        protocolSetting['security'] = data.scy
      } else {
        protocolSetting['security'] = 'auto'
      }
      if ('type' in data && data['type'] != 'none' && data['type'] != '' && data['type'] != 'auto') {
        // 'type in vmess is not supported.'
        continue
      }
      var host = ''
      if ('net' in data && data['net'] != 'tcp') {
        if (!['ws', 'quic'].includes(data['net'])) {
          // throw 'transport protocol not supported.'
          console.log(data['net'])
          continue
        }
        transport['protocol'] = data['net']
        if (data['net'] == 'ws') {
          if ('host' in data) {
            host = data['host']
            transportSetting['headers'] = {
              'Host': host
            }
          }
          if ('path' in data) {
            transportSetting['path'] = data['path']
          } else {
            transportSetting['path'] = '/'
          }
        }
      }
      if ('tls' in data && data['tls'] == 'tls') {
        transport['tls_type'] = 'tls'
        if ('sni' in data && data['sni'] != '') {
          transport['tls']['server_name'] = data['sni']
        } else if (host != '') {
          transport['tls']['server_name'] = host
        }
        if ('alpn' in data && data['alpn'] != '') {
          transport['tls']['alpn'] = data['alpn'].split(',')
        }
        if ('skip-cert-verify' in data && data['skip-cert-verify'] == true) {
          transport['tls']['insecure'] = true
        }
        if ('fp' in data && data['fp'] != '') {
          transport['tls']['utls'] = {
            'enabled': true,
            'fingerprint': data['fp']
          }
        }
      }
      address = data.add
      port = parseInt(data.port)
      tag = data.ps
    } else { //}}}
      var temp = parseInt(url.port)
      if (!isNaN(temp)) {
        port = temp
      }
      address = url.hostname
      var password = url.username
      if (address == '') {
        var temp = url.pathname.split('@')
        if (temp.length != 2) {
          continue
        }
        var addressAndPort = temp[1].split(':')
        if (temp.length != 2) {
          continue
        }
        address = addressAndPort[0]
        port = parseInt(addressAndPort[1])

        password = temp[0].replaceAll('/', '')
      } else {
        password = decodeURIComponent(password)
      }
      if (protocol == 'shadowsocks') {
        try {
          password = Base64.decode(password).split(':')
        } catch (e) {
          continue
        }
        var method = password[0]
        if (method == 'chacha20-poly1305') {
          method = 'chacha20-ietf-poly1305'
        }
        if (method == '') {
          continue
        }
        protocolSetting['method'] = method
        password = password[1]
      } else {
        transport['tls_type'] = 'tls'
        var sni = url.searchParams.get('sni')
        if (sni != undefined) {
          transport['tls']['server_name'] = sni
        }

        var type = url.searchParams.get('type')
        if (type == 'ws') {
          transport['protocol'] = 'ws'
          transportSetting['path'] = url.searchParams.get('path')
          var host = url.searchParams.get('host')
          if (host != undefined) {
            transportSetting['headers'] = {
              'Host': host
            }
          }
        }

        var insecure = url.searchParams.get('insecure')
        if (insecure == '1' || insecure == 'true') {
          transport['tls']['insecure'] = true
        }

        var obfs = url.searchParams.get('obfs')
        if (obfs != '' && obfs != undefined && obfs != 'none') {
          protocolSetting['obfs'] = {
            'type': obfs,
            'password': url.searchParams.get('obfs-password')
          }
        }
      }
      protocolSetting['password'] = password
      tag = decodeURIComponent(url.hash.substring(1))
    }
    if (transport['tls_type'] == 'tls') {
      transport['tls']['enabled'] = true
    }
    if (isNaN(port)) {
      continue
    }
    transport['address'] = address
    transport['port'] = port
    transport['setting'] = transportSetting

    var itemRes = {
      'protocol': protocol,
      'tag': tag,
      'setting': protocolSetting,
      'transport': transport
    }
    res.push(itemRes)
  }
  return res;
}
