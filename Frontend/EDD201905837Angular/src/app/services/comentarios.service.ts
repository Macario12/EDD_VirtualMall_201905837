import { Injectable } from '@angular/core';
import { Tienda } from  '../models/tienda.model';
import { HttpClient, HttpHandler, HttpHeaders } from '@angular/common/http'
import { Observable } from 'rxjs';
import {Global} from '../services/global.service'
@Injectable({
  providedIn: 'root'
})
export class ComentariosService {
  public url: string;

  constructor(public _http: HttpClient) {

    this.url =  Global.url;

  }

  getComentariosTienda(json:any):Observable<any> {
    let params = json
    let headers = new HttpHeaders().set('Content-Type','multipart/form-data');
    return this._http.post<any>(this.url+'getComentariosTienda', params,{headers:headers})
  }

  subirComentarioTienda(json:any):Observable<any> {
    let params = json
    let headers = new HttpHeaders().set('Content-Type','multipart/form-data');
      return this._http.post<any>(this.url+'subirComentarioTienda', params,{headers:headers})
  }
}
