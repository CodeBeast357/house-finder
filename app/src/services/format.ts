import numbro from "numbro";

numbro.registerLanguage({
  languageTag: "fr-CA",
  delimiters: {
    thousands: " ",
    decimal: ",",
  },
  abbreviations: {
    thousand: "k",
    million: "M",
    billion: "G",
    trillion: "T",
  },
  ordinal: (number) => {
    return number === 1 ? "er" : "ème";
  },
  spaceSeparated: true,
  currency: {
    symbol: "$",
    position: "postfix",
    code: "USD",
  },
  currencyFormat: {
    thousandSeparated: true,
    totalLength: 4,
    spaceSeparated: true,
    average: true,
  },
  formats: {
    fourDigits: {
      totalLength: 4,
      spaceSeparated: true,
      average: true,
    },
    fullWithTwoDecimals: {
      thousandSeparated: true,
      mantissa: 2,
    },
    fullWithTwoDecimalsNoCurrency: {
      mantissa: 2,
      thousandSeparated: true,
    },
    fullWithNoDecimals: {
      output: "currency",
      thousandSeparated: true,
      mantissa: 0,
    },
  },
});
numbro.setLanguage("fr-CA");
export const formatPrice = (price: number): string => {
  return numbro(price).formatCurrency({
    average: false,
    spaceSeparatedCurrency: true,
    thousandSeparated: true,
    currencyPosition: "postfix",
  });
};
export const formatRoundPrice = (price: number): string => {
  return numbro(price).formatCurrency({
    average: false,
    spaceSeparatedCurrency: true,
    thousandSeparated: true,
    forceAverage: "thousand",
    currencyPosition: "postfix",
  });
};
