/** Excel */
import * as xlsx from 'xlsx';

const { utils, writeFile, read } = xlsx;
const DEF_FILE_NAME = 'new-excel.xlsx';

function getHeaderRow(sheet: xlsx.WorkSheet) {
  const headers: string[] = [];
  const range = utils.decode_range(sheet['!ref'] as string);
  let C: number;
  const R = range.s.r;
  /* start in the first row */
  for (C = range.s.c; C <= range.e.c; ++C) {
    /* walk every column in the range */
    const cell = sheet[utils.encode_cell({ c: C, r: R })];
    /* find the cell in the first row */
    let hdr = `UNKNOWN ${C}`; // <-- replace with your desired default
    if (cell && cell.t) hdr = utils.format_cell(cell);
    headers.push(hdr);
  }
  return headers;
}

// 导入
export function importsExcel(file: { raw: Blob }) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();

    reader.onload = (e) => {
      // @ts-ignore
      const data = new Uint8Array(e.target.result);
      const workbook = read(data, { type: 'array' });
      //  console.log("workbook: ", workbook);

      // 将Excel 第一个sheet内容转为json格式
      const worksheet = workbook.Sheets[workbook.SheetNames[0]];
      const json = utils.sheet_to_json(worksheet);
      //   console.log("jsonExcel:", jsonExcel);
      const headers = getHeaderRow(worksheet);
      resolve({ tableData: json, headers });
    };

    reader.readAsArrayBuffer(file.raw);
  });
}

// 导出
export function aoaToSheetXlsx({
  data,
  header,
  filename = DEF_FILE_NAME,
  write2excelOpts = { bookType: 'xlsx' },
}) {
  const arrData = [...data];
  if (header) {
    arrData.unshift(header);
  }

  const worksheet = utils.aoa_to_sheet(arrData);
  /* add worksheet to workbook */
  const workbook = {
    SheetNames: [filename],
    Sheets: {
      [filename]: worksheet,
    },
  };
  /* output format determined by filename */
  // @ts-ignore
  writeFile(workbook, filename, write2excelOpts);
  /* at this point, out.xlsb will have been downloaded */
}
