import {Producto} from '../models/producto.model';
export class Tienda{
    constructor(
        public Nombre:string,
        public Descripcion:string,
        public Contacto:string,
        public Calificacion:string,
        public Logo:string,
        public Productos: Producto[],
        public Arbol64: string,
        public Raiz: any,
        public Pedidos64: string,
        public Fecha: string
        
    ){}
}