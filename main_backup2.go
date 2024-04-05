package main

//
//import (
//	"fmt"
//	"github.com/jung-kurt/gofpdf"
//)
//
////// Transaction represents a bank transaction
////type Transaction struct {
////	Date        string
////	Description string
////	Amount      float64
////}
////
////func main() {
////	// Create a new PDF document
////	pdf := gofpdf.New("P", "mm", "A4", "")
////
////	// Add a page
////	pdf.AddPage()
////
////	// Set font for header
////	pdf.SetFont("Arial", "B", 16)
////
////	// Calculate width of the page
////	width, _ := pdf.GetPageSize()
////
////	// Header text
////	headerText := "Bank Statement"
////	headerTextWidth := pdf.GetStringWidth(headerText)
////	headerTextX := (width - headerTextWidth) / 2
////	headerTextY := 20.0 // Position of header text
////
////	// Add header text
////	pdf.Text(headerTextX, headerTextY, headerText)
////
////	// Set font for statement details
////	pdf.SetFont("Arial", "", 12)
////
////	// Statement period
////	startDate := "2024-04-01"
////	endDate := "2024-04-30"
////	statementPeriod := fmt.Sprintf("Statement Period: %s - %s", startDate, endDate)
////	statementPeriodY := headerTextY + 15.0
////
////	// Add statement period
////	pdf.Text(10, statementPeriodY, statementPeriod)
////
////	// Add a line below the statement period
////	lineY := statementPeriodY + 10
////	pdf.Line(10, lineY, width-10, lineY)
////
////	// Set font for table header
////	pdf.SetFont("Arial", "B", 12)
////
////	// Create table header
////	columns := []string{"Date", "Description", "Amount"}
////	columnWidths := []float64{30, 100, 40}
////	//alignments := []string{"L", "L", "R"}
////
////	// Set top margin
////	pdf.SetY(lineY + 10)
////	// Add table header
////	for i, col := range columns {
////		pdf.CellFormat(columnWidths[i], 10, col, "", 0, "C", false, 0, "")
////	}
////	pdf.Ln(-1)
////
////	// Set font for table rows
////	pdf.SetFont("Arial", "", 12)
////
////	// Simulate bank transactions
////	transactions := []Transaction{
////		{Date: "2024-04-05", Description: "Salary", Amount: 2500.00},
////		{Date: "2024-04-10", Description: "Electricity Bill", Amount: -100.00},
////		{Date: "2024-04-15", Description: "Internet Bill", Amount: -50.00},
////		{Date: "2024-04-20", Description: "Groceries", Amount: -200.00},
////		{Date: "2024-04-25", Description: "Dinner", Amount: -40.00},
////	}
////	transactionDummy := Transaction{Date: "2024-04-05", Description: "Salary", Amount: 2500.00}
////	for i := 0; i < 20; i++ {
////		transactions = append(transactions, transactionDummy)
////	}
////
////	//Add table rows
////	for _, tx := range transactions {
////		pdf.CellFormat(columnWidths[0], 10, tx.Date, "", 0, "L", false, 0, "")
////		pdf.CellFormat(columnWidths[1], 10, tx.Description, "", 0, "L", false, 0, "")
////		pdf.CellFormat(columnWidths[2], 10, fmt.Sprintf("$%.2f", tx.Amount), "", 0, "R", false, 0, "")
////		pdf.Ln(-1)
////
////		// Check if we need to add a page break
////		if pdf.GetY() > 250 {
////			// Add a page break and start a new page
////			pdf.AddPage()
////			// Add table header
////			for i, col := range columns {
////				pdf.CellFormat(columnWidths[i], 10, col, "", 0, "C", false, 0, "")
////			}
////			//
////			//// Add some content to the new page
////			//addContent(pdf)
////			pdf.SetY(10)
////		}
////	}
////	// Save the PDF to a file
////	err := pdf.OutputFileAndClose("bank_statement.pdf")
////	if err != nil {
////		fmt.Println("Error:", err)
////		return
////	}
////
////	fmt.Println("Bank Statement PDF created successfully.")
////}
//
////type Transaction struct {
////	Date        string
////	Description string
////	Amount      float64
////}
////
////func main() {
////	// Create a new PDF
////	pdf := gofpdf.New("P", "mm", "A4", "")
////	// Set the font for the entire PDF
////	pdf.SetFont("Arial", "", 12)
////
////	// Get the width and height of the page
////	width, height := pdf.GetPageSize()
////
////	// Add a header to the PDF
////	headerTextY := 10.0
////	headerText := "Bank Statement"
////	pdf.Text(10, headerTextY, headerText)
////
////	// Statement period
////	startDate := "2024-04-01"
////	endDate := "2024-04-30"
////	statementPeriod := fmt.Sprintf("Statement Period: %s - %s", startDate, endDate)
////	statementPeriodY := headerTextY + 15.0
////
////	// Add statement period
////	pdf.Text(10, statementPeriodY, statementPeriod)
////
////	// Add a line below the statement period
////	lineY := statementPeriodY + 10
////	pdf.Line(10, lineY, width-10, lineY)
////
////	// Set font for table header
////	pdf.SetFont("Arial", "B", 12)
////
////	// Create table header
////	columns := []string{"Date", "Description", "Amount"}
////	columnWidths := []float64{30, 100, 40}
////
////	// Set top margin
////	pdf.SetY(lineY + 10)
////	// Add table header
////	for i, col := range columns {
////		pdf.CellFormat(columnWidths[i], 10, col, "", 0, "C", false, 0, "")
////	}
////	pdf.Ln(-1)
////
////	// Set font for table rows
////	pdf.SetFont("Arial", "", 12)
////
////	// Simulate bank transactions
////	transactions := []Transaction{
////		{Date: "2024-04-05", Description: "Salary", Amount: 2500.00},
////		{Date: "2024-04-10", Description: "Electricity Bill", Amount: -100.00},
////		{Date: "2024-04-15", Description: "Internet Bill", Amount: -50.00},
////		{Date: "2024-04-20", Description: "Groceries", Amount: -200.00},
////		{Date: "2024-04-25", Description: "Dinner", Amount: -40.00},
////	}
////	transactionDummy := Transaction{Date: "2024-04-05", Description: "Salary", Amount: 2500.00}
////	for i := 0; i < 20; i++ {
////		transactions = append(transactions, transactionDummy)
////	}
////
////	//Add table rows
////	for _, tx := range transactions {
////		pdf.CellFormat(columnWidths[0], 10, tx.Date, "", 0, "L", false, 0, "")
////		pdf.CellFormat(columnWidths[1], 10, tx.Description, "", 0, "L", false, 0, "")
////		pdf.CellFormat(columnWidths[2], 10, fmt.Sprintf("$%.2f", tx.Amount), "", 0, "R", false, 0, "")
////		pdf.Ln(-1)
////
////		// Check if we need to add a page break
////		if pdf.GetY() > 250 {
////			// Add a page break and start a new page
////			pdf.AddPage()
////			// Add table headerfor i, col := range columns {
////			pdf.CellFormat(columnWidths[i], 10, col, "", 0, "C", false, 0, "")
////		}
////		pdf.Ln(-1)
////		// Add some content to the new page
////		pdf.Text(10, 20, "This is the new page content.")
////		pdf.SetY(10)
////	}
////
////	// Save the PDF to a file
////	err := pdf.OutputFileAndClose("bank_statement.pdf")
////	if err != nil {
////		fmt.Println("Error:", err)
////		return
////	}
////
////	fmt.Println("Bank Statement PDF created successfully.")
////}
//
//type Transaction struct {
//	Date        string
//	Description string
//	Amount      float64
//}
//
//func main() {
//	// Create a new PDF
//	pdf := gofpdf.New("P", "mm", "A4", "")
//	// Set the font for the entire PDF
//	pdf.SetFont("Arial", "", 12)
//
//	// Get the width and height of the page
//	width, height := pdf.GetPageSize()
//
//	// Add a header to the PDF
//	headerTextY := 10.0
//	headerText := "Bank Statement"
//	pdf.Text(10, headerTextY, headerText)
//
//	// Statement period
//	startDate := "2024-04-01"
//	endDate := "2024-04-30"
//	statementPeriod := fmt.Sprintf("Statement Period: %s - %s", startDate, endDate)
//	statementPeriodY := headerTextY + 15.0
//
//	// Add statement period
//	pdf.Text(10, statementPeriodY, statementPeriod)
//
//	// Add a line below the statement period
//	lineY := statementPeriodY + 10
//	pdf.Line(10, lineY, width-10, lineY)
//
//	// Set font for table header
//	pdf.SetFont("Arial", "B", 12)
//
//	// Create table header
//	columns := []string{"Date", "Description", "Amount"}
//	columnWidths := []float64{30, 100, 40}
//
//	// Set top margin
//	pdf.SetY(lineY + 10)
//	// Add table header
//	for i, col := range columns {
//		pdf.CellFormat(columnWidths[i], 10, col, "", 0, "C", false, 0, "")
//	}
//	pdf.Ln(-1)
//
//	// Set font for table rows
//	pdf.SetFont("Arial", "", 12)
//
//	// Simulate bank transactions
//	transactions := []Transaction{
//		{Date: "2024-04-05", Description: "Salary", Amount: 2500.00},
//		{Date: "2024-04-10", Description: "Electricity Bill", Amount: -100.00},
//		{Date: "2024-04-15", Description: "Internet Bill", Amount: -50.00},
//		{Date: "2024-04-20", Description: "Groceries", Amount: -200.00},
//		{Date: "2024-04-25", Description: "Dinner", Amount: -40.00},
//	}
//	transactionDummy := Transaction{Date: "2024-04-05", Description: "Salary", Amount: 2500.00}
//	for i := 0; i < 50; i++ {
//		transactions = append(transactions, transactionDummy)
//	}
//
//	// Initialize row counter
//	rowCount := 0
//
//	//Add table rows
//	for _, tx := range transactions {
//		// Check if the current row is the last row that can fit on the page
//		if pdf.GetY()+10 > height-30 && rowCount > 0 {
//			// Add a page break and start a new page
//			pdf.AddPage()
//			// Add table header
//			for i, col := range columns {
//				pdf.CellFormat(columnWidths[i], 10, col, "", 0, "C", false, 0, "")
//			}
//			pdf.Ln(-1)
//			// Resetrow counter
//			rowCount = 0
//		}
//
//		// Add a new row to the table
//		pdf.CellFormat(columnWidths[0], 10, tx.Date, "", 0, "L", false, 0, "")
//		pdf.CellFormat(columnWidths[1], 10, tx.Description, "", 0, "L", false, 0, "")
//		pdf.CellFormat(columnWidths[2], 10, fmt.Sprintf("$%.2f", tx.Amount), "",
//			0, "R", false, 0, "")
//		pdf.Ln(-1)
//
//		// Increment row counter
//		rowCount++
//	}
//	// Save the PDF to a file
//	err := pdf.OutputFileAndClose("bank_statement.pdf")
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//
//	fmt.Println("Bank Statement PDF created successfully.")
//	//
//}
