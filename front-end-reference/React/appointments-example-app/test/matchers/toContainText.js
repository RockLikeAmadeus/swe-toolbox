export const toContainText = (received, expectedText) => ({
  pass: received.textContent.includes(expectedText),
});
