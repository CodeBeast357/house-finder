import { memoize } from "lodash/fp";
import { House, ProviderName } from "@/store/modules/house";

export const getHouseLogo = memoize((providerName: ProviderName): string => {
  return `assets/${providerName}.png`;
});

export const getGoogleMapsLink = memoize(
  ({ latitude, longitude }: House): string => {
    return `https://www.google.com/maps/place/${longitude},${latitude}`;
  }
);
