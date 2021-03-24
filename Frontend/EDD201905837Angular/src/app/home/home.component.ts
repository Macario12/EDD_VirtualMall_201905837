import { Component, OnInit, Inject } from '@angular/core';
import { Tienda } from '../models/tienda.model';
import { Producto } from '../models/producto.model';

import { DomSanitizer } from '@angular/platform-browser';
import { TiendasService } from '../services/tiendas.service';
import { ProductosService } from '../services/productos.service';

import {MatDialog, MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';




@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  public Tiendas: Tienda[];
  public currentInput;
  public status:string;
  selectedFle
  nombre
  imageSource;

  
  constructor(private _tiendasService: TiendasService, private _productosService: ProductosService, public dialogM: MatDialog
    ,private sanitizer: DomSanitizer) { }

  ngOnInit(): void {
    this.getStore();
  }

  getStore(){
    this._tiendasService.getTiendas().subscribe((tiendaAPI: Tienda[]) =>{

      this.Tiendas = tiendaAPI
      console.log(this.Tiendas)

    },error => console.error(error))
  }
  
  onFileSelected(event: any) {
    if(event.target.files.length > 0) 
     {
       console.log(event.target.files[0])
      this.selectedFle = event.target.files[0];
      this.nombre = this.selectedFle.name

      const reader = new FileReader();
      reader.onload = (e) => {
        const text = reader.result!.toString().trim();
        console.log(text)
        this._tiendasService.postTiendas(text).subscribe(
         Response => {
           console.log(Response)
           
         }
       )
       
        
      }
      reader.readAsText(this.selectedFle);
     }
   }

  postTiendas(){
    this._tiendasService.postTiendas(this.currentInput).subscribe(
      Response => {
        console.log(Response)
        if(Response){
          this.status = 'sucess';
        }
      }
    )
  }

  onFileSelectedInventario(event: any) {
    if(event.target.files.length > 0) 
     {
       console.log(event.target.files[0])
      this.selectedFle = event.target.files[0];
      this.nombre = this.selectedFle.name

      const reader = new FileReader();
      reader.onload = (e) => {
        const text = reader.result!.toString().trim();
        console.log(text)
        this._productosService.postProductos(text).subscribe(
         Response => {
           console.log(Response)
           
         }
       )
       
        
      }
      reader.readAsText(this.selectedFle);
     }
   }
   
   verArbolModal(imagen: string){
    this.imageSource = this.sanitizer.bypassSecurityTrustResourceUrl(`data:image/png;base64, ${imagen}`);
   }

   verProductos(productos: Producto[]){

      this._productosService.productosR = productos
      console.log(productos)
   }
}



