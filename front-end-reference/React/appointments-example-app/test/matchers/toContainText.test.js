import { toContainText } from "./toContainText";

describe("toContainText matcher", () => {
  it("returns pass is true when text is found in the given DOM element", () => {
    const domElement = { textContent: "text to find" };
    const result = toContainText(domElement, "text to find");
    expect(result.pass).toBe(true);
  });
  it("returns pass is false when the text is not found in the given DOM element", () => {
    const domElement = { textContent: "" };
    const result = toContainText(domElement, "text to find");
    expect(result.pass).toBe(false);
  });
});
