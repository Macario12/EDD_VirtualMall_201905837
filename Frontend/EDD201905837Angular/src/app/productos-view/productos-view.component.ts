import { Component, OnInit } from '@angular/core';
import { ProductosService } from '../services/productos.service';
import {Producto } from '../models/producto.model'

@Component({
  selector: 'app-productos-view',
  templateUrl: './productos-view.component.html',
  styleUrls: ['./productos-view.component.css']
})
export class ProductosViewComponent implements OnInit {

  constructor(public _productosService: ProductosService) { }

  public productR

  ngOnInit(): void {
    this.productR = this._productosService.productosR
  }

  
  detalleProducto(productoD: Producto){
    this._productosService.productoDetalle = productoD
  }


}
