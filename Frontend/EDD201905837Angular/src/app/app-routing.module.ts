import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CarritoDeComprasComponent } from './carrito-de-compras/carrito-de-compras.component';
import { HomeComponent } from './home/home.component';
import {ProductosViewComponent} from './productos-view/productos-view.component'
import { DialogComponent } from './dialog/dialog.component'
import { CalendarioComponent } from './calendario/calendario.component';

const routes: Routes = [
  {
    path: 'carrito',
    component: CarritoDeComprasComponent,
  },
  {
    path: 'home',
    component: HomeComponent,
  },
  {
    path: 'inventario',
    component: ProductosViewComponent,
  },
  {
    path: 'productoDetalles',
    component: DialogComponent,
  },
  {
    path: 'calendario',
    component: CalendarioComponent,
  },
  {
    path: '',
    component: HomeComponent,
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }