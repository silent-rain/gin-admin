import * as xlsx from 'xlsx';
const { utils, writeFile, read } = xlsx;
const DEF_FILE_NAME = 'new-excel.xlsx';
function getHeaderRow(sheet) {
    const headers = [];
    const range = utils.decode_range(sheet['!ref']);
    let C;
    const R = range.s.r;
    for (C = range.s.c; C <= range.e.c; ++C) {
        const cell = sheet[utils.encode_cell({ c: C, r: R })];
        let hdr = `UNKNOWN ${C}`;
        if (cell && cell.t)
            hdr = utils.format_cell(cell);
        headers.push(hdr);
    }
    return headers;
}
export function importsExcel(file) {
    return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onload = (e) => {
            const data = new Uint8Array(e.target.result);
            const workbook = read(data, { type: 'array' });
            const worksheet = workbook.Sheets[workbook.SheetNames[0]];
            const json = utils.sheet_to_json(worksheet);
            const headers = getHeaderRow(worksheet);
            resolve({ tableData: json, headers });
        };
        reader.readAsArrayBuffer(file.raw);
    });
}
export function aoaToSheetXlsx({ data, header, filename = DEF_FILE_NAME, write2excelOpts = { bookType: 'xlsx' }, }) {
    const arrData = [...data];
    if (header) {
        arrData.unshift(header);
    }
    const worksheet = utils.aoa_to_sheet(arrData);
    const workbook = {
        SheetNames: [filename],
        Sheets: {
            [filename]: worksheet,
        },
    };
    writeFile(workbook, filename, write2excelOpts);
}
