import { Injectable } from '@angular/core';
import { HttpClient, HttpHandler, HttpHeaders } from '@angular/common/http'
import { Observable } from 'rxjs';
import {Global} from '../services/global.service'
import { Usuario } from '../models/usuario.model';

@Injectable({
  providedIn: 'root'
})
export class UsuariosService {
  public url: string;
  public UsuarioA: Usuario;
  constructor(public _http: HttpClient) {

    this.url =  Global.url;

  }

  buscarUsuario(user: any): Observable<any>{
    let params = JSON.stringify(user)
    let headers = new HttpHeaders().set('Content-Type','multipart/form-data');
    return this._http.post(this.url+'buscarUsuario', params, {headers: headers});
  }

  crearUsuario(user: any): Observable<any>{
    let params = JSON.stringify(user)
    let headers = new HttpHeaders().set('Content-Type','multipart/form-data');
    return this._http.post(this.url+'crearUsuario', params, {headers: headers});
  }

  postUsuarios(users: any): Observable<any>{
    let params = users
    let headers = new HttpHeaders().set('Content-Type','multipart/form-data');

    return this._http.post(this.url+'usuarios',params,{headers:headers})
  }

  posrtllave(llave: any): Observable<any>{
    let params =  JSON.stringify(llave)
    let headers = new HttpHeaders().set('Content-Type','multipart/form-data');
    return this._http.post(this.url+'llave',params,{headers:headers});
  }

  getArbolS(): Observable<string>{
    return this._http.get<string>(this.url+'arbolS');
  }
  getArbolC(): Observable<string>{
    return this._http.get<string>(this.url+'arbolC');
  }
  getArbolCS(): Observable<string>{
    return this._http.get<string>(this.url+'arbolCS');
  }

  getGrafo(): Observable<string>{
    return this._http.get<string>(this.url+'grafo');
  }

  postGrafo(users: any): Observable<any>{
    let params = users
    let headers = new HttpHeaders().set('Content-Type','multipart/form-data');

    return this._http.post(this.url+'Cgrafo',params,{headers:headers})
  }
}
