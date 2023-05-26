import {Routes} from "@angular/router";
import {PageComponent} from "./page.component";

export const routes:Routes = [
  {
    path:'',
    component: PageComponent,
    children:[
      {path:'',redirectTo:'movie',pathMatch:'full'},
      {
        path:'movie',
        loadChildren:()=>import('./movie/routes').then((m)=>m.routes)
      }
    ]
  }
]
