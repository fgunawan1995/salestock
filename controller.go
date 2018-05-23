package main
import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
    "os"
    "encoding/csv"
)
//Controller

//Items
func PostItem(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var Item Items
	c.Bind(&Item)
	if Item.SKU != "" && Item.NamaItem != "" && Item.JumlahSekarang > -1 {
		db.Create(&Item)
		c.JSON(201, gin.H{"success": Item})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}
func GetItems(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var Items []Items
	db.Find(&Items)
	c.JSON(200, Items)
}
func GetItem(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var Item Items
	db.First(&Item, id)
	if Item.Id != 0 {
		c.JSON(200, Item)
	} else {
		c.JSON(404, gin.H{"error": "Item not found"})
	}
}
func UpdateItem(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var Item Items
	db.First(&Item, id)
	if Item.SKU != "" && Item.NamaItem != "" && Item.JumlahSekarang > -1 {
		if Item.Id != 0 {
			var newItem Items
			c.Bind(&newItem)
			result := Items{
				Id:        Item.Id,
				SKU: newItem.SKU,
				NamaItem:  newItem.NamaItem,
				JumlahSekarang:  newItem.JumlahSekarang,
			}
			db.Save(&result)
			c.JSON(200, gin.H{"success": result})
		} else {
			c.JSON(404, gin.H{"error": "Item not found"})
		}
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}
func DeleteItem(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var Item Items
	db.First(&Item, id)
	if Item.Id != 0 {
		db.Delete(&Item)
		c.JSON(200, gin.H{"success": "Item #" + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "Item not found"})
	}
}
func OptionsItem(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

//ItemsIn
func PostItemIn(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var ItemIn ItemsIn
	c.Bind(&ItemIn)
	if ItemIn.Waktu != "" && ItemIn.SKU != "" && ItemIn.NamaBarang != "" && ItemIn.JumlahPemesanan > -1 && ItemIn.JumlahDiterima > -1 && ItemIn.HargaBeli > -1 && ItemIn.Total > -1 && ItemIn.NomorKwitansi != "" && ItemIn.Catatan != ""{
		db.Create(&ItemIn)
		c.JSON(201, gin.H{"success": ItemIn})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}
func GetItemsIn(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var ItemsIn []ItemsIn
	db.Find(&ItemsIn)
	c.JSON(200, ItemsIn)
}
func GetItemIn(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var ItemIn ItemsIn
	db.First(&ItemIn, id)
	if ItemIn.Id != 0 {
		c.JSON(200, ItemIn)
	} else {
		c.JSON(404, gin.H{"error": "ItemIn not found"})
	}
}
func UpdateItemIn(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var ItemIn ItemsIn
	db.First(&ItemIn, id)
	if ItemIn.Waktu != "" && ItemIn.SKU != "" && ItemIn.NamaBarang != "" && ItemIn.JumlahPemesanan > -1 && ItemIn.JumlahDiterima > -1 && ItemIn.HargaBeli > -1 && ItemIn.Total > -1 && ItemIn.NomorKwitansi != "" && ItemIn.Catatan != ""{
		if ItemIn.Id != 0 {
			var newItem ItemsIn
			c.Bind(&newItem)
			result := ItemsIn{
				Id:        ItemIn.Id,
				Waktu: newItem.Waktu,
				SKU: newItem.SKU,
				NamaBarang:  newItem.NamaBarang,
				JumlahPemesanan:  newItem.JumlahPemesanan,
				JumlahDiterima:  newItem.JumlahDiterima,
				HargaBeli:  newItem.HargaBeli,
				Total:  newItem.Total,
				NomorKwitansi:  newItem.NomorKwitansi,
				Catatan:  newItem.Catatan,
			}
			db.Save(&result)
			c.JSON(200, gin.H{"success": result})
		} else {
			c.JSON(404, gin.H{"error": "ItemIn not found"})
		}
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}
func DeleteItemIn(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var ItemIn ItemsIn
	db.First(&ItemIn, id)
	if ItemIn.Id != 0 {
		db.Delete(&ItemIn)
		c.JSON(200, gin.H{"success": "ItemIn #" + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "ItemIn not found"})
	}
}
func OptionsItemIn(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

//ItemsOut
func PostItemOut(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var ItemOut ItemsOut
	c.Bind(&ItemOut)
	if ItemOut.Waktu != "" && ItemOut.SKU != "" && ItemOut.NamaBarang != "" && ItemOut.JumlahKeluar > -1 && ItemOut.HargaJual > -1 && ItemOut.Total > -1 && ItemOut.Catatan != ""{
		db.Create(&ItemOut)
		c.JSON(201, gin.H{"success": ItemOut})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}
func GetItemsOut(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var ItemsOut []ItemsOut
	db.Find(&ItemsOut)
	c.JSON(200, ItemsOut)
}
func GetItemOut(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var ItemOut ItemsOut
	db.First(&ItemOut, id)
	if ItemOut.Id != 0 {
		c.JSON(200, ItemOut)
	} else {
		c.JSON(404, gin.H{"error": "ItemOut not found"})
	}
}
func UpdateItemOut(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var ItemOut ItemsOut
	db.First(&ItemOut, id)
	if ItemOut.Waktu != "" && ItemOut.SKU != "" && ItemOut.NamaBarang != "" && ItemOut.JumlahKeluar > -1 && ItemOut.HargaJual > -1 && ItemOut.Total > -1 && ItemOut.Catatan != ""{
		if ItemOut.Id != 0 {
			var newItem ItemsOut
			c.Bind(&newItem)
			result := ItemsOut{
				Id:        ItemOut.Id,
				Waktu: newItem.Waktu,
				SKU: newItem.SKU,
				NamaBarang:  newItem.NamaBarang,
				JumlahKeluar:  newItem.JumlahKeluar,
				HargaJual:  newItem.HargaJual,
				Total:  newItem.Total,
				Catatan:  newItem.Catatan,
			}
			db.Save(&result)
			c.JSON(200, gin.H{"success": result})
		} else {
			c.JSON(404, gin.H{"error": "ItemOut not found"})
		}
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}
func DeleteItemOut(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var ItemOut ItemsOut
	db.First(&ItemOut, id)
	if ItemOut.Id != 0 {
		db.Delete(&ItemOut)
		c.JSON(200, gin.H{"success": "ItemOut #" + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "ItemOut not found"})
	}
}
func OptionsItemOut(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

//Reporting
func GetItemsReport(c *gin.Context) {
	db := InitDb()
	var ItemsReport []ItemsReport
	db.Table("items").Select("items.sku as sku, items.nama_item as nama_item, (AVG(items_ins.harga_beli)*(items.jumlah_sekarang + SUM(items_ins.jumlah_diterima) - (SELECT SUM(jumlah_keluar) FROM items_outs where items_outs.sku=items.sku))) as total, (items.jumlah_sekarang + SUM(items_ins.jumlah_diterima) - (SELECT SUM(jumlah_keluar) FROM items_outs where items_outs.sku=items.sku)) as jumlah, AVG(items_ins.harga_beli) as rata").Joins("left join items_ins on items.sku = items_ins.sku").Group("items.sku").Scan(&ItemsReport)
	c.JSON(200, ItemsReport)
}
func GetSalesReport(c *gin.Context) {
	db := InitDb()
	var SalesReport []SalesReport
	db.Table("items_outs").Select("catatan as id_pesanan, waktu, sku, nama_barang, jumlah_keluar as jumlah, harga_jual, total, (SELECT AVG(items_ins.harga_beli) FROM items left join items_ins on items.sku = items_ins.sku  where items.sku=items_outs.sku GROUP BY items.sku) as harga_beli, (harga_jual-(SELECT AVG(items_ins.harga_beli) FROM items left join items_ins on items.sku = items_ins.sku  where items.sku=items_outs.sku GROUP BY items.sku))*jumlah_keluar as laba").Scan(&SalesReport)
	c.JSON(200, SalesReport)
}

//Convert to CSV
func GetCsvItems(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var Items []Items
	db.Find(&Items)
	
	csvdatafile, err := os.Create("items.csv")

    if err != nil {
       fmt.Println(err)
    }
    defer csvdatafile.Close()

    writer := csv.NewWriter(csvdatafile)

	var record []string
        record = append(record, "SKU")
        record = append(record, "Nama Item")
		record = append(record, "Jumlah Sekarang")
        writer.Write(record)
		
    for _, worker := range Items {
        var record []string
        record = append(record, worker.SKU)
        record = append(record, worker.NamaItem)
		record = append(record, strconv.Itoa(worker.JumlahSekarang))
        writer.Write(record)
    }
    writer.Flush()
	c.JSON(200, Items)
}
func GetCsvItemsIn(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var ItemsIn []ItemsIn
	db.Find(&ItemsIn)
	
	csvdatafile, err := os.Create("itemsin.csv")

    if err != nil {
       fmt.Println(err)
    }
    defer csvdatafile.Close()

    writer := csv.NewWriter(csvdatafile)

	var record []string
		record = append(record, "Waktu")
        record = append(record, "SKU")
        record = append(record, "Nama Barang")
		record = append(record, "Jumlah Pemesanan")
		record = append(record, "Jumlah Diterima")
		record = append(record, "Harga Beli")
		record = append(record, "Total")
		record = append(record, "Nomor Kwitansi")
		record = append(record, "Catatan")
        writer.Write(record)
		
    for _, worker := range ItemsIn {
        var record []string
		record = append(record, worker.Waktu)
        record = append(record, worker.SKU)
        record = append(record, worker.NamaBarang)
		record = append(record, strconv.Itoa(worker.JumlahPemesanan))
		record = append(record, strconv.Itoa(worker.JumlahDiterima))
		record = append(record, strconv.Itoa(worker.HargaBeli))
		record = append(record, strconv.Itoa(worker.Total))
		record = append(record, worker.NomorKwitansi)
        record = append(record, worker.Catatan)
        writer.Write(record)
    }
    writer.Flush()
	c.JSON(200, ItemsIn)
}
func GetCsvItemsOut(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var ItemsOut []ItemsOut
	db.Find(&ItemsOut)
	
	csvdatafile, err := os.Create("itemsout.csv")

    if err != nil {
       fmt.Println(err)
    }
    defer csvdatafile.Close()

    writer := csv.NewWriter(csvdatafile)

	var record []string
		record = append(record, "Waktu")
        record = append(record, "SKU")
        record = append(record, "Nama Barang")
		record = append(record, "Jumlah Keluar")
		record = append(record, "Harga Jual")
		record = append(record, "Total")
		record = append(record, "Catatan")
        writer.Write(record)
		
    for _, worker := range ItemsOut {
        var record []string
		record = append(record, worker.Waktu)
        record = append(record, worker.SKU)
        record = append(record, worker.NamaBarang)
		record = append(record, strconv.Itoa(worker.JumlahKeluar))
		record = append(record, strconv.Itoa(worker.HargaJual))
		record = append(record, strconv.Itoa(worker.Total))
        record = append(record, worker.Catatan)
        writer.Write(record)
    }
    writer.Flush()
	c.JSON(200, ItemsOut)
}
func GetCsvItemsReport(c *gin.Context) {
	db := InitDb()
	var ItemsReport []ItemsReport
	db.Table("items").Select("items.sku as sku, items.nama_item as nama_item, (AVG(items_ins.harga_beli)*(items.jumlah_sekarang + SUM(items_ins.jumlah_diterima) - (SELECT SUM(jumlah_keluar) FROM items_outs where items_outs.sku=items.sku))) as total, (items.jumlah_sekarang + SUM(items_ins.jumlah_diterima) - (SELECT SUM(jumlah_keluar) FROM items_outs where items_outs.sku=items.sku)) as jumlah, AVG(items_ins.harga_beli) as rata").Joins("left join items_ins on items.sku = items_ins.sku").Group("items.sku").Scan(&ItemsReport)
	
	csvdatafile, err := os.Create("itemsreport.csv")

    if err != nil {
       fmt.Println(err)
    }
    defer csvdatafile.Close()

    writer := csv.NewWriter(csvdatafile)

	var record []string
        record = append(record, "SKU")
        record = append(record, "Nama Item")
		record = append(record, "Jumlah")
		record = append(record, "Rata-Rata Harga Beli")
		record = append(record, "Total")
        writer.Write(record)
		
    for _, worker := range ItemsReport {
        var record []string
        record = append(record, worker.SKU)
        record = append(record, worker.NamaItem)
		record = append(record, strconv.Itoa(worker.Jumlah))
		record = append(record, strconv.Itoa(worker.Rata))
		record = append(record, strconv.Itoa(worker.Total))
        writer.Write(record)
    }
    writer.Flush()
	c.JSON(200, ItemsReport)
}
func GetCsvSalesReport(c *gin.Context) {
	db := InitDb()
	var SalesReport []SalesReport
	db.Table("items_outs").Select("catatan as id_pesanan, waktu, sku, nama_barang, jumlah_keluar as jumlah, harga_jual, total, (SELECT AVG(items_ins.harga_beli) FROM items left join items_ins on items.sku = items_ins.sku GROUP BY items.sku) as harga_beli, (harga_jual-(SELECT AVG(items_ins.harga_beli) FROM items left join items_ins on items.sku = items_ins.sku GROUP BY items.sku))*jumlah_keluar as laba").Scan(&SalesReport)
	
	csvdatafile, err := os.Create("salesreport.csv")

    if err != nil {
       fmt.Println(err)
    }
    defer csvdatafile.Close()

    writer := csv.NewWriter(csvdatafile)

	var record []string
        record = append(record, "ID Pesanan")
        record = append(record, "Waktu")
        record = append(record, "SKU")
        record = append(record, "Nama Barang")
		record = append(record, "Jumlah")
		record = append(record, "Harga Beli")
		record = append(record, "Laba")
        writer.Write(record)
		
    for _, worker := range SalesReport {
        var record []string
        record = append(record, worker.IdPesanan)
		record = append(record, worker.Waktu)
		record = append(record, worker.SKU)
        record = append(record, worker.NamaBarang)
		record = append(record, strconv.Itoa(worker.Jumlah))
		record = append(record, strconv.Itoa(worker.HargaBeli))
		record = append(record, strconv.Itoa(worker.Laba))
        writer.Write(record)
    }
    writer.Flush()
	c.JSON(200, SalesReport)
}