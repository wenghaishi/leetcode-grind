const a =
  "Regular follow up meeting with Amgen to discuss potential MOU support package for Singapore subsidiary Amgen Singapore Manufacturing (“ASM”) and the company’s future business plans in Singapore.";

const b =
  'Regular follow up meeting with Amgen to discuss potential MOU support package for Singapore subsidiary Amgen Singapore Manufacturing ("ASM") and the company\'s future business plans in Singapore.';

const normalize = (str) =>
  str.replace(/’/g, "'").replace(/“/g, '"').replace(/”/g, '"'); // Convert curly quotes to normal quotes

console.log(normalize(a).includes(normalize(b))); // true ✅

const c =
  "executives and AM aligned on potential focus areas for DIAL 3.0, in view of company's push towards Business AI, specifically on AI in specific domain areas (e.g., healthcare, financial services.";

const d =
  "video info, while company value-adds in the software/analysis layer (by translating the info captured into tangible analysis and recommendations) Future plans for DIAL 3. 0 executives and AM aligned on potential focus areas for DIAL 3. 0, in view of company’s push towards Business AI, specifically on AI in specific domain areas (e. g. , healthcare, financial services) - On healthcare, company shared increased appetite among customers in singapore & region (e. g. , IHH) noting strong potential and appetite to address higher value-added XM use cases (e. g. , patient and physician engagement) AM shared broadly on our goals and efforts for NAIS2. 0, including the push towards accelerating lead demand /";

console.log(d.includes(c));
