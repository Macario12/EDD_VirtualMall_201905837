import { Injectable } from '@angular/core';
import { Tienda } from  '../models/tienda.model';
import { HttpClient, HttpHandler, HttpHeaders } from '@angular/common/http'
import { Observable } from 'rxjs';
import {Global} from '../services/global.service'

@Injectable({
  providedIn: 'root'
})
export class TiendasService {
  public url: string;

  constructor(public _http: HttpClient) {

    this.url =  Global.url;

  }

  getTiendas(): Observable<Tienda[]>{
    return this._http.get<Tienda[]>(this.url+'tiendascargadas');
  }

  postTiendas(tiendas: any): Observable<any>{
    let params = tiendas
    let headers = new HttpHeaders().set('Content-Type','multipart/form-data');

    return this._http.post(this.url+'cargartienda',params,{headers:headers})
  }
  
}
