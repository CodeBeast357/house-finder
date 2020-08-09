import mapboxgl from "mapbox-gl";

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
    id: "schools",
    type: "symbol",
    source: "schools",
    "source-layer": "schools",
    layout: {
      "icon-image": "school",
      "icon-allow-overlap": true,
    },
  });
};
