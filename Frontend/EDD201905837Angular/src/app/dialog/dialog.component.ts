import { Component, OnInit } from '@angular/core';
import { ProductosService } from '../services/productos.service';
import {Producto } from '../models/producto.model'
import { ProductoC } from  '../models/productoCom.model';
@Component({
  selector: 'app-dialog',
  templateUrl: './dialog.component.html',
  styleUrls: ['./dialog.component.css']
})
export class DialogComponent implements OnInit {
  public producotDetalle
  public cantidadElegida
  public vectorProdcutoD: Producto[]
  public status: boolean
  constructor(public _productosService: ProductosService) { 
    this.vectorProdcutoD = new Array()
  }

  ngOnInit(): void {
    this.producotDetalle = this._productosService.productoDetalle
    
  }

  AgregaraCarrito(productoAcarrito: Producto){
    this.validarCantidad()
    if(this.status == false){
      this.producotDetalle.Cantidad = this.cantidadElegida
      this._productosService.carritoCompras.push(productoAcarrito)
      console.log(this._productosService.carritoCompras)
    }
  }

  validarCantidad(){
    if(this.cantidadElegida > this.producotDetalle.Cantidad){
        this.status = true;
    }else{
        this.status = false;
    }
  }
}
