package main

//
//import (
//	"fmt"
//	"github.com/jung-kurt/gofpdf"
//	"math"
//	"strconv"
//)
//
////func main() {
////	//pdf := gofpdf.New("P", "mm", "A4", "")
////	//pdf.AddPage()
////	//pdf.SetFont("Arial", "B", 16)
////	////pdf.Text(40, 10, "Hello, world")
////	////pdf.Image("./sample.png", 56, 40, 100, 0, false, "", 0, "")
////	//// Set font
////	//pdf.SetFont("Arial", "B", 16)
////	//
////	//// Add text
////	//pdf.CellFormat(10, 10, "Hello, World!", "1", 0, "M", false, 0, "")
////	//
////	//err := pdf.OutputFileAndClose("./file.pdf")
////	//if err != nil {
////	//	log.Println("ERROR", err.Error())
////	//}
////	//Create a new PDF document
////	pdf := gofpdf.New("P", "mm", "A4", "")
////
////	// Add a page
////	pdf.AddPage()
////
////	// Set font for header
////	pdf.SetFont("Arial", "B", 16)
////
////	// Calculate width and height of the page
////	width, _ := pdf.GetPageSize()
////
////	// Header text
////	headerText := "Invoice"
////	headerTextWidth := pdf.GetStringWidth(headerText)
////	headerTextX := (width - headerTextWidth) / 2
////	headerTextY := 20 // Position of header text
////
////	// Add header text
////	pdf.Text(headerTextX, float64(headerTextY), headerText)
////
////	// Add a line below the header
////	lineY := headerTextY + 10
////	pdf.Line(10, float64(lineY), width-10, float64(lineY))
////
////	// Set font for table
////	pdf.SetFont("Arial", "", 12)
////
////	// Create a simple table
////	data := [][]string{
////		{"Name", "Quantity", "Price"},
////		{"Apple", "10", "$2.50"},
////		{"Banana", "5", "$1.20"},
////		{"Orange", "8", "$1.80"},
////	}
////
////	// Set column widths
////	columnWidths := []float64{50, 30, 30}
////	alignments := []string{"L", "C", "C"}
////
////	// Loop through data and add to table
////	for _, row := range data {
////		for i, str := range row {
////			pdf.CellFormat(columnWidths[i], 10, str, "", 0, alignments[i], false, 0, "")
////		}
////		pdf.Ln(-1) // Move to the next line
////	}
////
////	// Save the PDF to a file
////	err := pdf.OutputFileAndClose("output.pdf")
////	if err != nil {
////		fmt.Println("Error:", err)
////		return
////	}
////
////	fmt.Println("PDF created successfully.")
////}
//
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
////	// Calculate width and height of the page
////	width, _ := pdf.GetPageSize()
////
////	// Header text
////	headerText := "Invoice"
////	headerTextWidth := pdf.GetStringWidth(headerText)
////	headerTextX := (width - headerTextWidth) / 2
////	headerTextY := 20.0 // Position of header text
////
////	// Add header text
////	pdf.Text(headerTextX, headerTextY, headerText)
////
////	// Set font for invoice details
////	pdf.SetFont("Arial", "", 12)
////
////	// Invoice details
////	invoiceDetails := []string{
////		"Invoice Number: 12345",
////		"Date: 2024-04-25",
////		"Customer: John Doe",
////	}
////
////	// Set Y position for invoice details
////	invoiceDetailsY := headerTextY + 50
////
////	// Add invoice details
////	for _, detail := range invoiceDetails {
////		pdf.Text(10, invoiceDetailsY, detail)
////		invoiceDetailsY += 10
////	}
////
////	// Add a line below the invoice details
////	lineY := invoiceDetailsY + 10
////	pdf.Line(10, lineY, width-10, lineY)
////
////	// Set font for table header
////	pdf.SetFont("Arial", "B", 12)
////
////	// Create table header
////	columns := []string{"Item", "Quantity", "Price", "Total"}
////	columnWidths := []float64{60, 30, 40, 40}
////	alignments := []string{"L", "C", "R", "R"}
////
////	// Add table header
////	for i, col := range columns {
////		pdf.CellFormat(columnWidths[i], 50, col, "", 0, "C", false, 0, "")
////	}
////	pdf.Ln(-1)
////
////	// Set font for table rows
////	pdf.SetFont("Arial", "", 12)
////
////	// Table data
////	data := [][]string{
////		{"Item 1", "2", "$10.00", "$20.00"},
////		{"Item 2", "1", "$15.00", "$15.00"},
////		{"Item 3", "3", "$8.00", "$24.00"},
////	}
////
////	// Add table rows
////	for _, row := range data {
////		for i, str := range row {
////			pdf.CellFormat(columnWidths[i], 10, str, "", 0, alignments[i], false, 0, "")
////		}
////		pdf.Ln(-1)
////	}
////
////	// Calculate total
////	var total float64
////	for _, row := range data {
////		price, _ := strconv.ParseFloat(row[2][1:], 64) // Remove "$" from price
////		quantity, _ := strconv.Atoi(row[1])
////		total += price * float64(quantity)
////	}
////
////	// Add total
////	pdf.CellFormat(width-20, 10, "Total", "", 0, "R", false, 0, "")
////	pdf.CellFormat(40, 10, fmt.Sprintf("$%.2f", total), "", 0, "R", false, 0, "")
////
////	// Save the PDF to a file
////	err := pdf.OutputFileAndClose("invoice.pdf")
////	if err != nil {
////		fmt.Println("Error:", err)
////		return
////	}
////
////	fmt.Println("Invoice PDF created successfully.")
////}
//
//func main() {
//	marginX := 10.0
//	marginY := 20.0
//	pdf := gofpdf.New("P", "mm", "A4", "")
//	pdf.SetMargins(marginX, marginY, marginX)
//	pdf.AddPage()
//
//	pdf.ImageOptions("assets/logo.png", 0, 0, 65, 25, false, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "")
//
//	pdf.SetFont("Arial", "B", 16)
//	_, lineHeight := pdf.GetFontSize()
//	currentY := pdf.GetY() + lineHeight
//	pdf.SetY(currentY)
//	pdf.Cell(40, 10, "Company Name")
//
//	lineHt := 10.0
//	const colNumber = 5
//	header := [colNumber]string{"No", "Description", "Quantity", "Unit Price ($)", "Price ($)"}
//	colWidth := [colNumber]float64{10.0, 75.0, 25.0, 40.0, 40.0}
//
//	// Headers
//	pdf.SetFontStyle("B")
//	pdf.SetFillColor(200, 200, 200)
//	for colJ := 0; colJ < colNumber; colJ++ {
//		pdf.CellFormat(colWidth[colJ], lineHt, header[colJ], "1", 0, "CM", true, 0, "")
//	}
//
//	pdf.Ln(-1)
//
//	// Table data
//	pdf.SetFontStyle("")
//	subtotal := 0.0
//
//	for rowJ := 0; rowJ < len(data); rowJ++ {
//		val := data[rowJ]
//		if len(val) == 3 {
//			// Column 1: Unit
//			// Column 2: Description
//			// Column 3: Price per unit
//			unit, _ := strconv.Atoi(val[0])
//			desc := val[1]
//			pricePerUnit, _ := strconv.ParseFloat(val[2], 64)
//			pricePerUnit = math.Round(pricePerUnit*100) / 100
//			totalPrice := float64(unit) * pricePerUnit
//			subtotal += totalPrice
//
//			pdf.CellFormat(colWidth[0], lineHt, fmt.Sprintf("%d", rowJ+1), "1", 0, "CM", false, 0, "")
//			pdf.CellFormat(colWidth[1], lineHt, desc, "1", 0, "LM", false, 0, "")
//			pdf.CellFormat(colWidth[2], lineHt, fmt.Sprintf("%d", unit), "1", 0, "CM", false, 0, "")
//			pdf.CellFormat(colWidth[3], lineHt, fmt.Sprintf("%.2f", pricePerUnit), "1", 0, "CM", false, 0, "")
//			pdf.CellFormat(colWidth[4], lineHt, fmt.Sprintf("%.2f", totalPrice), "1", 0, "CM", false, 0, "")
//			pdf.Ln(-1)
//		}
//	}
//
//	// Calculate the subtotal
//	pdf.SetFontStyle("B")
//	leftIndent := 0.0
//	for i := 0; i < 3; i++ {
//		leftIndent += colWidth[i]
//	}
//	pdf.SetX(marginX + leftIndent)
//	pdf.CellFormat(colWidth[3], lineHt, "Subtotal", "1", 0, "CM", false, 0, "")
//	pdf.CellFormat(colWidth[4], lineHt, fmt.Sprintf("%.2f", subtotal), "1", 0, "CM", false, 0, "")
//	pdf.Ln(-1)
//
//	taxAmount := math.Round(subtotal*float64(taxPercent)) / 100
//	pdf.SetX(marginX + leftIndent)
//	pdf.CellFormat(colWidth[3], lineHt, "Tax Amount", "1", 0, "CM", false, 0, "")
//	pdf.CellFormat(colWidth[4], lineHt, fmt.Sprintf("%.2f", taxAmount), "1", 0, "CM", false, 0, "")
//	pdf.Ln(-1)
//
//	grandTotal := subtotal + taxAmount
//	pdf.SetX(marginX + leftIndent)
//	pdf.CellFormat(colWidth[3], lineHt, "Grand total", "1", 0, "CM", false, 0, "")
//	pdf.CellFormat(colWidth[4], lineHt, fmt.Sprintf("%.2f", grandTotal), "1", 0, "CM", false, 0, "")
//	pdf.Ln(-1)
//}
