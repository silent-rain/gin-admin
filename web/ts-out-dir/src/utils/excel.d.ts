export declare function importsExcel(file: {
    raw: Blob;
}): Promise<unknown>;
export declare function aoaToSheetXlsx({ data, header, filename, write2excelOpts, }: {
    data: any;
    header: any;
    filename?: string | undefined;
    write2excelOpts?: {
        bookType: string;
    } | undefined;
}): void;
