import { Component} from '@angular/core';
import {BrowserModule} from "@angular/platform-browser";
import {FormsModule} from "@angular/forms";
import {HttpClientModule} from "@angular/common/http";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
import {RouterModule} from "@angular/router";

@Component({
  selector: 'app-root',
  template: '<router-outlet></router-outlet>',
  styles: [':host{display: flex;flex: 1}'],
  standalone:true,
  imports: [
    RouterModule,
    BrowserModule,
    FormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
  ],
})
export class AppComponent {

}
