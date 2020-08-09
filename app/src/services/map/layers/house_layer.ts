import mapboxgl, { EventData, LngLatLike } from "mapbox-gl";

import { curry, flow, head, get } from "lodash/fp";
import { House } from "@/store/modules/house";
import { store } from "@/store/store";
import { initComponent } from "@/services/init_vue";
import HousePopup from "@/components/map/HousePopup.vue";

const HOUSE_TYPE_CONFIG = {
  normal: {
    layerName: "normal-houses",
    sourceName: "normal-houses",
    fileName: "marker.png",
    imageName: "house-marker",
    layout: {},
  },
  star: {
    layerName: "star-houses",
    sourceName: "star-houses",
    fileName: "star.png",
    imageName: "house-star",
    layout: {
      "icon-size": 0.1,
    },
  },
};

const getAddHouseLayer = (houseType: "normal" | "star") => (
  initialHouseList: House[],
  mapInstance: mapboxgl.Map
) => {
  const iconConfig = HOUSE_TYPE_CONFIG[houseType];
  // eslint-disable-next-line @typescript-eslint/no-var-requires
  const marker = require(`@/assets/${iconConfig.fileName}`);
  return new Promise((resolve) => {
    mapInstance.loadImage(marker, (error: Error, image: ImageData) => {
      if (error) throw error;

      mapInstance.addImage(iconConfig.imageName, image, { sdf: true });
      mapInstance.addSource(iconConfig.sourceName, {
        type: "vector",
        tiles: ["http://localhost/tiles/maps/houses/{z}/{x}/{y}.vector.pbf?"],
      });

      mapInstance.addLayer({
        id: iconConfig.layerName,
        type: "symbol",
        source: iconConfig.sourceName,
        "source-layer": "houses",
        layout: {
          "icon-image": iconConfig.imageName,
          "icon-allow-overlap": true,
          ...iconConfig.layout,
        },
        paint: {
          "icon-color": "#5dbcd2",
          "icon-halo-color": "yellow",
          "icon-halo-width": [
            "case",
            ["boolean", ["feature-state", "selected"], false],
            1,
            0,
          ],
        },
      });
      updateLayer(houseType)(initialHouseList, mapInstance);
      mapInstance.once("sourcedata", () => {
        resolve();
      });
    });
  });
};

export const addHouseLayer = getAddHouseLayer("normal");
export const addStarHouseLayer = getAddHouseLayer("star");

const getFeatures = curry(
  (mapInstance: mapboxgl.Map, event: mapboxgl.EventData) =>
    mapInstance.queryRenderedFeatures(event.point, {
      layers: [
        HOUSE_TYPE_CONFIG.normal.layerName,
        HOUSE_TYPE_CONFIG.star.layerName,
      ],
    })
);

const getFirstFeatureHouse = flow(
  head,
  (feature: mapboxgl.MapboxGeoJSONFeature) => {
    return feature
      ? {
          id: feature.id,
        }
      : null;
  }
);

const renderMapPopup = (
  mapInstance: mapboxgl.Map,
  mapPopupCoordinates: LngLatLike
) => {
  const mapPopupContainer = document.createElement("div");
  const mapPopupChild = document.createElement("div");
  mapPopupContainer.setAttribute("id", "house-popup");
  mapPopupContainer.appendChild(mapPopupChild);

  const popup = new mapboxgl.Marker(mapPopupContainer, {
    offset: [-30, 0],
    anchor: "right",
  })
    .setLngLat([0, 0])
    .addTo(mapInstance);

  popup.setLngLat(mapPopupCoordinates);
  initComponent(HousePopup, mapPopupChild);
};

export const destroyMapPopup = (): void => {
  const elem = document.querySelector("#house-popup");
  elem && elem.parentNode.removeChild(elem);
};

export const addOnHouseClickHandler = (
  onSelectedHouse: ({ houseId: number }) => void,
  onDeselectHouse: () => void
) => (mapInstance: mapboxgl.Map): void => {
  mapInstance.on("click", (event: EventData) => {
    if (event.originalEvent.target.localName !== "canvas") return;
    flow(getFeatures(mapInstance), getFirstFeatureHouse, (mapFeature) => {
      if (mapFeature) {
        const { id } = mapFeature;
        onSelectedHouse({ houseId: id });
        destroyMapPopup();
        renderMapPopup(mapInstance, event.lngLat);
      } else {
        destroyMapPopup();
        onDeselectHouse();
      }
    })(event);
  });
};

const updateSelected = (
  mapInstance: mapboxgl.Map,
  houseId: number,
  selected: boolean
): void => {
  if (houseId) {
    const oldHouse: House = store.getters["house/findHouse"](houseId);
    const oldLayerToUpdate =
      HOUSE_TYPE_CONFIG[oldHouse.isInSweetSpot ? "star" : "normal"].layerName;
    mapInstance.setFeatureState(
      {
        source: oldLayerToUpdate,
        sourceLayer: "houses",
        id: houseId,
      },
      {
        selected,
      }
    );
  }
};

export const updateFeatureState = (
  mapInstance: mapboxgl.Map,
  houseId: number,
  oldHouseId: number
): void => {
  updateSelected(mapInstance, houseId, true);
  updateSelected(mapInstance, oldHouseId, false);
};

const updateLayer = (houseType: "normal" | "star") => (
  houseList: House[],
  mapInstance: mapboxgl.Map
): void => {
  const idList = houseList.map(get("id"));
  const filter = ["in", "house_id", ...idList];
  mapInstance.setFilter(HOUSE_TYPE_CONFIG[houseType].layerName, filter);
};

export const updateHouseLayer = updateLayer("normal");
export const updateStarHouseLayer = updateLayer("star");
