import Vue from "vue";

import { store } from "@/store/store";
import vuetify from "@/plugins/vuetify";

export const initComponent = (
  component: typeof Vue.prototype.$createElement,
  selector: Element | string
): Vue => {
  return new Vue({
    render: (h) => h(component),
    vuetify,
    store,
  }).$mount(selector);
};
