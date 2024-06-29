<template>
  <div>
    <el-form label-width="150px">
      <el-row :gutter="5">
        <el-col :xs="24" :sm="12" :md="8" :lg="8" :xl="8">
          <el-tooltip :disabled="uif.showToolTip" placement="top">
            <div slot="content">出口转发, 小心本地回环</div>
            <el-form-item label="detour">
              <out_seletor
                :placeholder="Translator({ cn: '选填', en: 'Optional' })"
                :outbound="outbound_obj.dial.detour"
                :isDetour="true"
              />
            </el-form-item>
          </el-tooltip>
        </el-col>
      </el-row>
    </el-form>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
import { InitSetting } from "@/utils";
import out_seletor from "@/uif_views/outbounds/my_servers/out_seletor.vue";

export default {
  name: "dial",
  props: ["outbound_obj"],
  components: { out_seletor },
  data() {
    return {};
  },
  created() {
    var setting = {
      detour: { id: [], tag: "" },
    };
    if (!("dial" in this.outbound_obj)) {
      this.outbound_obj.dial = {};
    }
    this.outbound_obj.dial = InitSetting(this.outbound_obj.dial, setting);
    console.log(this.outbound_obj.dial);
  },
  computed: {
    ...mapState(["uif"]),
  },

  methods: {
    ...mapActions({}),
    Translator(i) {
      return i[this.uif.config.lang];
    },
  },
};
</script>
