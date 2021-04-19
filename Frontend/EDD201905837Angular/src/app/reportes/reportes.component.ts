import { Component, OnInit } from '@angular/core';
import { DomSanitizer } from '@angular/platform-browser';
import { UsuariosService } from '../services/usuarios.service';
@Component({
  selector: 'app-reportes',
  templateUrl: './reportes.component.html',
  styleUrls: ['./reportes.component.css']
})
export class ReportesComponent implements OnInit {
  public imageSource;
  public arbol64: string
  constructor(private sanitizer: DomSanitizer,private _usuarioService: UsuariosService) { }

  ngOnInit(): void {
  }
  getArbolS(){
    this._usuarioService.getArbolS().subscribe((arbolBase: string) =>{

      this.arbol64 = arbolBase
      this.verArbolModal(this.arbol64)
      console.log(this.arbol64)

    },error => console.error(error))
  }

  getArbolC(){
    this._usuarioService.getArbolC().subscribe((arbolBase: string) =>{

      this.arbol64 = arbolBase
      this.verArbolModal(this.arbol64)
      console.log(this.arbol64)

    },error => console.error(error))
  }

  getArbolCS(){
    this._usuarioService.getArbolCS().subscribe((arbolBase: string) =>{

      this.arbol64 = arbolBase
      this.verArbolModal(this.arbol64)
      console.log(this.arbol64)

    },error => console.error(error))
  }

  verArbolModal(imagen: string){
    this.imageSource = this.sanitizer.bypassSecurityTrustResourceUrl(`data:image/png;base64, ${imagen}`);
   }
}
