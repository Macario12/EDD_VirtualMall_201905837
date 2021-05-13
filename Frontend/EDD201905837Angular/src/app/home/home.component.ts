import { Component, OnInit, Inject } from '@angular/core';
import { Tienda } from '../models/tienda.model';
import { Producto } from '../models/producto.model';

import { Registro, TablaHash } from '../models/Tabla.model';
import { DomSanitizer } from '@angular/platform-browser';
import { TiendasService } from '../services/tiendas.service';
import { ProductosService } from '../services/productos.service';
import { PedidosService } from '../services/pedidos.service';
import { ComentariosService } from '../services/comentarios.service';

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
  public IdtablahashHija: any
  selectedFle
  nombre
  imageSource;
  Fecha;
  imageSource1;
  tienda: any
  public comentario:string = ""
  usuario:any
  preuba: number[]  = [1,2,3,4,6]
  registros: Registro[]  = []
  
  constructor(private _tiendasService: TiendasService, private _usuariosService: UsuariosService, private _productosService: ProductosService, public dialogM: MatDialog
    ,private sanitizer: DomSanitizer, private _PedidosService: PedidosService, private peticionesService: ComentariosService) {
      this.llave = new LLave("")
      this.comentario = ""
     }

  ngOnInit(): void {
    this.getStore();
  }
  tiendaU(tienda :Tienda){
    this.tienda = tienda
    console.log(this.tienda)
  }

  guardarId(id:string){
    this.IdtablahashHija = id
  }
  getComentarios(){
    this.usuario = this._usuariosService.UsuarioA
      var requestBody = {
        Tienda: this.tienda
      };
      this.peticionesService.getComentariosTienda(JSON.stringify(requestBody)).subscribe((res:Registro[]) => {
        this.registros = res
        console.log(this.registros)
      })
  }
  makeComment(idTabla:string) {
      alert(this.comentario)
      alert(idTabla)
      var requestBody = {
        Tienda: this.tienda,
        Usuario: this._usuariosService.UsuarioA,
        IdTabla: this.IdtablahashHija,
        Comentario: this.comentario,
      };
      
      console.log(requestBody)
      this.peticionesService.subirComentarioTienda(JSON.stringify(requestBody)).subscribe((res:Registro[]) => {
        this.registros = res
        console.log(this.registros)
      })
    
     this. IdtablahashHija = ""  
    this.getComentarios()
  }

  mostrar(registro:any) {
    alert(registro.Comentario)
    return registro.Comentario
  }
  createRange(number:number){
    var items: number[] = [];
    for(var i = 1; i <= number; i++){
       items.push(i);
    }
    return items;
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
        console.log(Response)
        if(Response){
          this.status = 'sucess';
        }
      }
    )
  }

  getStore(){
    this._tiendasService.getTiendas().subscribe((tiendaAPI: Tienda[]) =>{

      this.Tiendas = tiendaAPI
      

    },error => console.error(error))
    console.log(this.Tiendas)
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



