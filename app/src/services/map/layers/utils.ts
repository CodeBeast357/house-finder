import { SCHOOL_LAYER_ID } from "./schools";

const toggleLayer = (layerId: string) => (
  shouldShow: boolean,
  mapInstance: mapboxgl.Map
): void => {
  mapInstance.setLayoutProperty(
    layerId,
    "visibility",
    shouldShow ? "visible" : "none"
  );
};

export const toggleSchoolsLayer = toggleLayer(SCHOOL_LAYER_ID);
