def reversed_excel_sheet_column_table(columnNumber: str):
  mapper = {chr(i + 64): i for i in range(1,27)}
  for index, element in enumerate(columnNumber[::-1]):
      value = (26 ** index) * (mapper[element])
      output = output + value
  return output
