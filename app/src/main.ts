import Vue from "vue";
import "material-design-icons-iconfont/dist/material-design-icons.css";
import "mapbox-gl/dist/mapbox-gl.css";

import App from "./App.vue";
import { initComponent } from "./services/init_vue";

Vue.config.productionTip = false;

initComponent(App, "#app");
