package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

// Transaction represents a bank transaction
type Transaction struct {
	Date        string
	Description string
	Amount      float64
}

type Statement struct {
	TxType          string `json:"txType"`
	TxRef           string `json:"txRef"`
	TxMsg           string `json:"txMsg"`
	EndBalance      string `json:"endBalance"`
	ReferenceNumber string `json:"referenceNumber"`
	TxDate          string `json:"txDate"`
	TxAmt           string `json:"txAmt"`
	Time            string `json:"time"`
	PageCount       string `json:"pageCount"`
	Index           int    `json:"index"`
}

func main() {
	// Create a new PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add a page
	pdf.AddPage()

	// Set font for header
	pdf.SetFont("Arial", "B", 16)

	// Calculate width of the page
	width, height := pdf.GetPageSize()

	// Header text
	headerText := "Mutasi Rekening"
	headerTextWidth := pdf.GetStringWidth(headerText)
	headerTextX := (width - headerTextWidth) / 2
	headerTextY := 20.0 // Position of header text

	// Add header text
	pdf.Text(headerTextX, headerTextY, headerText)

	// Set font for statement details
	pdf.SetFont("Arial", "", 12)

	//Nomor rekening
	accountNumber := "7300002228"
	accountNumberField := fmt.Sprintf("Nomor Rekening: %s", accountNumber)
	accountNumberFieldY := headerTextY + 15.0
	// Add Nomor rekening
	pdf.Text(10, accountNumberFieldY, accountNumberField)

	//Nama
	name := "YOGI SAPUTRA"
	nameField := fmt.Sprintf("Nama: %s", name)
	nameFieldY := accountNumberFieldY + 5.0
	// Add Nomor rekening
	pdf.Text(10, nameFieldY, nameField)

	//Jenis Produk
	jenisProduk := "GIRO WADIAH PERUSAHAAN"
	jenisProdukField := fmt.Sprintf("Jenis Produk: %s", jenisProduk)
	jenisProdukFieldY := nameFieldY + 5.0
	// Add Nomor rekening
	pdf.Text(10, jenisProdukFieldY, jenisProdukField)

	//Mata Uang
	mataUang := "IDR"
	mataUangField := fmt.Sprintf("Mata Uang: %s", mataUang)
	mataUangFieldY := jenisProdukFieldY + 5.0
	// Add Nomor rekening
	pdf.Text(10, mataUangFieldY, mataUangField)

	// Statement period / Periode Transaksi
	startDate := "2024-04-01"
	endDate := "2024-04-30"
	statementPeriod := fmt.Sprintf("Periode Transaksi: %s - %s", startDate, endDate)
	statementPeriodY := mataUangFieldY + 5.0
	pdf.Text(10, statementPeriodY, statementPeriod)

	// SaldoAwal
	saldoAwal := "Rp. 933.038.548.23"
	saldoAwalField := fmt.Sprintf("Saldo Awal: %s", saldoAwal)
	saldoAwalFieldY := statementPeriodY + 5.0
	// Add statement period
	pdf.Text(10, saldoAwalFieldY, saldoAwalField)

	// SaldoAkhir
	saldoAkhir := "Rp. 936.459.833.72"
	saldoAkhirField := fmt.Sprintf("Saldo Akhir: %s", saldoAkhir)
	saldoAkhirFieldY := saldoAwalFieldY + 5.0
	// Add statement period
	pdf.Text(10, saldoAkhirFieldY, saldoAkhirField)
	//
	//// Add a line below the statement period
	//lineY := statementPeriodY + 10
	//pdf.Line(10, lineY, width-10, lineY)

	// Set font for table header
	pdf.SetFont("Arial", "B", 12)

	// Create table header
	columns := []string{"Date", "Description", "Amount"}
	columnWidths := []float64{30, 100, 40}
	//alignments := []string{"L", "L", "R"}

	// Set top margin
	pdf.SetY(saldoAkhirFieldY + 10)
	// Add table header
	for i, col := range columns {
		pdf.CellFormat(columnWidths[i], 10, col, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Simulate bank transactions
	transactions := []Transaction{
		{Date: "2024-04-05", Description: "Salary", Amount: 2500.00},
		{Date: "2024-04-10", Description: "Electricity Bill", Amount: -100.00},
		{Date: "2024-04-15", Description: "Internet Bill", Amount: -50.00},
		{Date: "2024-04-20", Description: "Groceries", Amount: -200.00},
		{Date: "2024-04-25", Description: "Dinner", Amount: -40.00},
	}
	transactionDummy := Transaction{Date: "2024-04-05", Description: "Salary", Amount: 2500.00}
	for i := 0; i < 20; i++ {
		transactions = append(transactions, transactionDummy)
	}

	////Add table rows
	//for _, tx := range transactions {
	//	pdf.CellFormat(columnWidths[0], 10, tx.Date, "", 0, "L", false, 0, "")
	//	pdf.CellFormat(columnWidths[1], 10, tx.Description, "", 0, "L", false, 0, "")
	//	pdf.CellFormat(columnWidths[2], 10, fmt.Sprintf("$%.2f", tx.Amount), "", 0, "R", false, 0, "")
	//	pdf.Ln(-1)
	//
	//	// Check if we need to add a page break
	//	if pdf.GetY() > 250 {
	//		// Add a page break and start a new page
	//		pdf.AddPage()
	//		// Set font for table Header
	//		// Add table header
	//		for i, col := range columns {
	//			pdf.CellFormat(columnWidths[i], 10, col, "", 0, "C", false, 0, "")
	//		}
	//		pdf.Ln(-1)
	//
	//	}
	//}
	// Initialize row counter
	rowCount := 0

	//Add table rows
	for _, tx := range transactions {
		// Check if the current row is the last row that can fit on the page
		if pdf.GetY()+10 > height-30 && rowCount > 0 {
			// Add a page break and start a new page
			pdf.AddPage()
			// Set font for table rows
			pdf.SetFont("Arial", "B", 12)
			// Add table header
			for i, col := range columns {
				pdf.CellFormat(columnWidths[i], 10, col, "1", 0, "C", false, 0, "")
			}
			pdf.Ln(-1)
			// Resetrow counter
			rowCount = 0
		}
		// Set font for table rows
		pdf.SetFont("Arial", "", 12)
		// Add a new row to the table
		pdf.CellFormat(columnWidths[0], 10, tx.Date, "1", 0, "L", false, 0, "")
		pdf.CellFormat(columnWidths[1], 10, tx.Description, "1", 0, "L", false, 0, "")
		pdf.CellFormat(columnWidths[2], 10, fmt.Sprintf("$%.2f", tx.Amount), "1",
			0, "R", false, 0, "")
		pdf.Ln(-1)

		// Increment row counter
		rowCount++
	}
	// Save the PDF to a file
	err := pdf.OutputFileAndClose("bank_statement.pdf")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Bank Statement PDF created successfully.")
}
