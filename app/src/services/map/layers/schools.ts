import mapboxgl from "mapbox-gl";

export const SCHOOL_LAYER_ID = "schools";

export const addSchoolLayer = (mapInstance: mapboxgl.Map): void => {
  mapInstance.addImage(
    "school",
    <HTMLImageElement>document.querySelector("#school-icon")
  );
  mapInstance.addSource("schools", {
    type: "vector",
    tiles: ["http://localhost/tiles/maps/schools/{z}/{x}/{y}.vector.pbf?"],
  });

  mapInstance.addLayer({
    id: SCHOOL_LAYER_ID,
    type: "symbol",
    source: "schools",
    "source-layer": "schools",
    layout: {
      "icon-image": "school",
      "icon-allow-overlap": true,
      visibility: "none",
    },
  });
};
