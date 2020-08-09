import mapboxgl from "mapbox-gl";

import { BoundingBox } from "@/store/modules/map";

export const initMapBox = ({
  x1,
  y1,
  x2,
  y2,
}: BoundingBox): Promise<mapboxgl.Map> => {
  mapboxgl.accessToken = process.env.VUE_APP_MAPBOX_TOKEN;

  const mapInstance = new mapboxgl.Map({
    container: "map-container",
    style: "mapbox://styles/mapbox/streets-v11",
    bounds: [
      [x1, y1],
      [x2, y2],
    ],
    interactive: true,
  });

  mapInstance.addControl(new mapboxgl.NavigationControl(), "bottom-right");

  return new Promise((resolve) => {
    mapInstance.on("load", () => {
      mapInstance.resize();
      resolve(mapInstance);
    });
  });
};
