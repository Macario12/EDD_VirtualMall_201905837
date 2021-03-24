import { Injectable } from '@angular/core';
import { Producto } from  '../models/producto.model';
import { HttpClient, HttpHandler, HttpHeaders } from '@angular/common/http'
import { Observable } from 'rxjs';
import {Global} from '../services/global.service'
@Injectable({
  providedIn: 'root'
})
export class ProductosService {

  public url: string;
  public productosR: Producto[]
  public carritoCompras: Producto[]
  
  public productoDetalle: Producto
  constructor(public _http: HttpClient) {
    this.carritoCompras = new Array()
    this.url =  Global.url;

  }

  postProductos(productos: any): Observable<any>{
    let params = productos
    let headers = new HttpHeaders().set('Content-Type','multipart/form-data');

    return this._http.post(this.url+'agregarINV',params,{headers:headers})
  }


}
