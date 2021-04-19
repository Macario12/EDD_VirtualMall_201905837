import { Component, OnInit, Inject } from '@angular/core';
import { Tienda } from '../models/tienda.model';
import { Producto } from '../models/producto.model';

import { DomSanitizer } from '@angular/platform-browser';
import { TiendasService } from '../services/tiendas.service';
import { ProductosService } from '../services/productos.service';
import { PedidosService } from '../services/pedidos.service';

import {MatDialog, MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';
import { UsuariosService } from '../services/usuarios.service';
import { LLave } from '../models/llave.model';




@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  public Tiendas: Tienda[];
  public currentInput;
  public status:string;
  public llave: LLave;
  selectedFle
  nombre
  imageSource;
 Fecha;
  imageSource1;
  
  constructor(private _tiendasService: TiendasService, private _usuariosService: UsuariosService, private _productosService: ProductosService, public dialogM: MatDialog
    ,private sanitizer: DomSanitizer, private _PedidosService: PedidosService) {
      this.llave = new LLave("")
     }

  ngOnInit(): void {
    this.getStore();
  }

  validarUsuario(): boolean{
    if(this._usuariosService.UsuarioA.Cuenta == "Admin"){
      return true
    }else{
      return false
    }
  }

  llavenviar(){
    this._usuariosService.posrtllave(this.llave).subscribe(
      Response => {
      console.log(this.llave)
        console.log(Response)
        
      }
    )
  }

  getStore(){
    this._tiendasService.getTiendas().subscribe((tiendaAPI: Tienda[]) =>{

      this.Tiendas = tiendaAPI
      console.log(this.Tiendas)

    },error => console.error(error))
  }
  
  onFileSelectedGrafo(event: any) {
    if(event.target.files.length > 0) 
     {
       console.log(event.target.files[0])
      this.selectedFle = event.target.files[0];
      this.nombre = this.selectedFle.name

      const reader = new FileReader();
      reader.onload = (e) => {
        const text = reader.result!.toString().trim();
        console.log("hey")
        console.log(text)
        this._usuariosService.postGrafo(text).subscribe(
         Response => {
           console.log(Response)
           
         }
       )
       
        
      }
      reader.readAsText(this.selectedFle);
     }
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
        console.log("hey")
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

   onFileSelectedUsuarios(event: any) {
    if(event.target.files.length > 0) 
     {
       console.log(event.target.files[0])
      this.selectedFle = event.target.files[0];
      this.nombre = this.selectedFle.name

      const reader = new FileReader();
      reader.onload = (e) => {
        const text = reader.result!.toString().trim();
        console.log("hey")
        console.log(text)
        this._usuariosService.postUsuarios(text).subscribe(
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

   verMatrizModal(imagen: string, fecha: string){
     this.Fecha = fecha
    this.imageSource1 = this.sanitizer.bypassSecurityTrustResourceUrl(`data:image/png;base64, ${imagen}`);
   }

   verProductos(productos: Producto[], riaz: any){

      this._productosService.productosR = productos
      this._productosService.riaz = riaz
      console.log(productos)
      console.log(riaz)
   }


   onFileSelectedPedido(event: any) {
    if(event.target.files.length > 0) 
     {
       console.log(event.target.files[0])
       this.selectedFle = event.target.files[0];
       this.nombre = this.selectedFle.name

      const reader = new FileReader();
      reader.onload = (e) => {
        const text = reader.result!.toString().trim();
        console.log(text)
          this._PedidosService.postPedidos(text).subscribe(
          Response => {
            console.log(Response)
            
          }
       )
       
        
      }
      reader.readAsText(this.selectedFle);
     }
   }
}



