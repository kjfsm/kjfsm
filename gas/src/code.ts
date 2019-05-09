function myFunction() {
  const sheetID = PropertiesService.getScriptProperties().getProperty("SHEET_ID");
  if (sheetID == null) {
    Logger.log("sheetIDが取得できません");
    throw new Error("sheetIDが取得できません");
  }
  const sheet = SpreadsheetApp.openById(sheetID);
  Logger.log(sheet.getName());
  console.log(sheet.getName());

}
