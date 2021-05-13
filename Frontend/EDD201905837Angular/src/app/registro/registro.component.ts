import { Component, OnInit, ViewChild, Input } from '@angular/core';
import {Router, ActivatedRoute, Params, RouterLink} from "@angular/router";
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Usuario } from '../models/usuario.model';
import { UsuariosService } from '../services/usuarios.service';
import * as CryptoJS from 'crypto-js'
@Component({
  selector: 'app-registro',
  templateUrl: './registro.component.html',
  styleUrls: ['./registro.component.css']
})
export class RegistroComponent implements OnInit {
  public usuario : Usuario
  public identity
  public UsuarioAdmin = "Admin"
  public secret = 'SECRET';
  
  public status:string;
  constructor(private _usuarioservice: UsuariosService,private _router: Router) { 
    this.usuario = new Usuario(0,"","","","Usuario");
  }

  ngOnInit(): void {
  }

  buscarUsuario(){
    this.usuario.Password = CryptoJS.SHA256(this.usuario.Password.trim()).toString();
    console.log(this.usuario)
    this._usuarioservice.buscarUsuario(this.usuario).subscribe(
      Respose =>{
        this.identity = Respose
        console.log(this.identity)
        if(!this.identity){
          this.status = 'error';
          console.log("No estas registrado");
          
        }else{
          console.log("Existe")
          localStorage.setItem('identity', JSON.stringify(this.identity));
          this._usuarioservice.UsuarioA = Respose
          console.log(this._usuarioservice.UsuarioA)
          this._router.navigate(['/home']);
        }
      },
      error =>{
        var errorMensage = <any> error;
        console.log(errorMensage);
        if(errorMensage != null){
          this.status = 'error';
        }
      }
    );
    console.log(this._usuarioservice.UsuarioA)
  }

  crearUsuario(){
    this.usuario.Password = CryptoJS.SHA256(this.usuario.Password.trim()).toString();
    this._usuarioservice.crearUsuario(this.usuario).subscribe(
      Response => {
        console.log(Response)
      }
    )
  }
}
