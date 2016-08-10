abstract class ExportUtil {

    private static function getExcelObject($headArr,$data){

        if(!is_array($data)){
            return null;
        }
        $objPHPExcel = new PHPExcel();
        $letters     = ['A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z'];
        $pointer = -1;
        for ($i = 0, $l = count($headArr); $i < $l; ++$i) {
            $p = $i % 26;
            if ($p == 0) ++$pointer;
            $letters[] = $letters[$pointer] . $letters[$p];
        }
        for($i=0,$l=count($headArr);$i<$l;++$i){
            $objPHPExcel->setActiveSheetIndex(0)->setCellValue($letters[$i].'1', $headArr[$i]);
        }
        $column = 2;
        $objActSheet = $objPHPExcel->getActiveSheet();
        foreach($data as $key => $rows){
            $i = 0;
            foreach($rows as $keyName=>$value){
                $objActSheet->setCellValueExplicit($letters[$i++].$column, $value, PHPExcel_Cell_DataType::TYPE_STRING);
            }
            $column++;
        }
        $objPHPExcel->getActiveSheet()->setTitle('Simple');
        $objPHPExcel->setActiveSheetIndex(0);

        return $objPHPExcel;
    }

    public static function saveToExcel($filename,$headArr,$data){

        if(empty($filename)){
            return false;
        }
        $objPHPExcel = self::getExcelObject($headArr,$data);
        $objWriter = PHPExcel_IOFactory::createWriter($objPHPExcel, 'Excel2007');

        try{
            $objWriter->save($filename);
            return true;
        }catch (PHPExcel_Writer_Exception $e){
            return false;
        }
    }

    public static function excelExport($fileName,$headArr,$data){

        if(empty($fileName)){
            return false;
        }
        $objPHPExcel = self::getExcelObject($headArr,$data);
        header('Content-Type: application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;charset=utf-8');
        header("Content-Disposition: attachment; filename=\"$fileName\"");
        header('Cache-Control: max-age=0');
        $objWriter = PHPExcel_IOFactory::createWriter($objPHPExcel, 'Excel2007');
        $objWriter->save('php://output');
    }

    public static function toDownload($file,$header=null,$name=null){

        if(null === $name){
            $name = pathinfo($file,PATHINFO_BASENAME);
        }
        if(is_array($header)){
            foreach($header as $row){
                header($row);
            }
        }else{
            header('Content-type: application/octet-stream');
            header('Content-Disposition: attachment; filename="'.$name.'"');
        }
        ob_clean();
        flush();
        readfile($file);
    }
}
