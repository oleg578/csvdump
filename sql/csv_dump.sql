-- Variables
SET @file_path := '/tmp/dump/employee.csv';
SET @table_name := 'employee';
SET @database_name := 'hr';

-- Drop temporary table if exists
DROP TEMPORARY TABLE IF EXISTS temp_table;

-- Create a temporary table with headers
CREATE TEMPORARY TABLE temp_table
SELECT COLUMN_NAME
FROM INFORMATION_SCHEMA.COLUMNS
WHERE TABLE_NAME = @table_name
  AND TABLE_SCHEMA = @database_name
ORDER BY ORDINAL_POSITION;

-- Dynamically create the SELECT statement with headers
SET @headers_query := (SELECT GROUP_CONCAT(CONCAT('"', COLUMN_NAME, '"') SEPARATOR ', ')
                       FROM temp_table);

SET @data_query := (SELECT GROUP_CONCAT(COLUMN_NAME SEPARATOR ', ')
                    FROM temp_table);

SET @full_query := CONCAT(
        'SELECT ', @headers_query,
        ' UNION ALL SELECT ', @data_query, ' FROM ', @table_name,
            ' INTO OUTFILE ''', @file_path,
        ''' FIELDS TERMINATED BY '','' ENCLOSED BY ''"'' LINES TERMINATED BY ''\n'''
                   );

-- Execute the full query to export data with headers
PREPARE stmt FROM @full_query;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;