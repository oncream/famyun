// import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';
//
// import { AppModule } from './app/app.module';
//
//
// platformBrowserDynamic().bootstrapModule(AppModule)
//   .catch(err => console.error(err));


import {bootstrapApplication} from "@angular/platform-browser";
import {AppComponent} from "./app/app.component";
import {provideRouter, RouterModule, Routes} from "@angular/router";
import {importProvidersFrom} from "@angular/core";
import {registerLocaleData} from "@angular/common";
import {zh_CN} from "ng-zorro-antd/i18n";

const routes: Routes = [
  {path:'',pathMatch:'full',redirectTo:'page'},
  {path:'page',
    loadChildren:() => import('./app/page/routes').then((m)=>m.routes)
  }
];

registerLocaleData(zh_CN)

bootstrapApplication(AppComponent,{
  providers:[
    provideRouter(routes)
  ]
}).catch(err=>console.log(err))



