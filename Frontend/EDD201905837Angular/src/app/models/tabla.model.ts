import { NumberSymbol } from "@angular/common"
import { Usuario } from "./usuario.model"

export class TablaHash {
    Id: string
    IndentLevel: number
    Registros: Registro[]
    Capacidad: number
    Carga: NumberSymbol
    constructor(Id: string, IndentLevel: number, Registros: Registro[], Capacidad: number, Carga: number) {
        this.Id = Id
        this.IndentLevel = IndentLevel
        this.Registros = Registros
        this.Capacidad = Capacidad
        this.Carga = Carga
    }   
}

export class Registro {
    IdContenedor: string
    Usuario: Usuario
    IdentLevel: number
    Comentario: string
    TablaHash: TablaHash
    constructor(IdContenedor: string, Usuario: Usuario, IdentLevel: number, Comentario: string, TablaHash: TablaHash) {
        this.IdContenedor = IdContenedor
        this.Usuario = Usuario
        this.IdentLevel  = IdentLevel
        this.Comentario = Comentario
        this.TablaHash = TablaHash
    }
}