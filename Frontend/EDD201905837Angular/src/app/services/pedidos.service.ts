import { Injectable } from '@angular/core';

import { Pedido } from  '../models/pedido.model';
import { HttpClient, HttpHandler, HttpHeaders } from '@angular/common/http'
import { Observable } from 'rxjs';
import {Global} from '../services/global.service'
@Injectable({
  providedIn: 'root'
})
export class PedidosService {
  public url
  constructor(public _http: HttpClient) {
    this.url =  Global.url;

  }
  postPedidos(pedidos: any): Observable<any>{
    let params = pedidos
    let headers = new HttpHeaders().set('Content-Type','multipart/form-data');

    return this._http.post(this.url+'cargarPedidos',params,{headers:headers})
  }

}
