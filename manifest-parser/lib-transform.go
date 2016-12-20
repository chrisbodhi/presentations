function appendPdf(val) {                   // HL
  return val + '.pdf';                      // HL
}                                           // HL

function removeNonAlphanumeric(val) {       // HL
  return val.replace(/\D/g, '');            // HL
}                                           // HL

function encodeAmpersand(val) {
  return val.replace(/&/g, '&amp;');
}

function decodeAmpersand(val) {
  return val.replace(/&amp;/g, '&');
}

function trimWhitespace(val) {
  return val.trim();
}
