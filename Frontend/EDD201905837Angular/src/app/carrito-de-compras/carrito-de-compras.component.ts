import { Component, OnInit } from '@angular/core';
import { ProductosService } from '../services/productos.service';
import {Producto } from '../models/producto.model'

@Component({
  selector: 'app-carrito-de-compras',
  templateUrl: './carrito-de-compras.component.html',
  styleUrls: ['./carrito-de-compras.component.css']
})
export class CarritoDeComprasComponent implements OnInit {
  public ACarrito
  public _total: number
  constructor(public _productosService: ProductosService) {
    
    this.ACarrito = this._productosService.carritoCompras
   }

  ngOnInit(): void {
    this._total = 0
    this.ACarrito = this._productosService.carritoCompras
    console.log(this._productosService.carritoCompras)
    for (let index = 0; index < this.ACarrito.length; index++) {
      this._total = this._total + this.ACarrito[index].Cantidad*this.ACarrito[index].Precio;
      
      
    }
    console.log(this._total)
  }

  
  removerItemProducto(productoEliminar: Producto){
     var i = this._productosService.carritoCompras.indexOf(productoEliminar)
     i !== -1 && this._productosService.carritoCompras.splice(i,1);
     this.ngOnInit()
  }
}
