import {Component} from '@angular/core';
import {Router, RouterModule} from "@angular/router";
import {CommonModule} from "@angular/common";
import {NzLayoutModule} from "ng-zorro-antd/layout";
import {NzIconModule} from "ng-zorro-antd/icon";
import {NzDropDownModule} from "ng-zorro-antd/dropdown";

@Component({
  selector: 'app-page',
  templateUrl: './page.component.html',
  styleUrls: ['./page.component.scss'],
  standalone:true,
  imports:[
    RouterModule,
    CommonModule,
    NzLayoutModule,
    NzIconModule,
    NzDropDownModule,
  ]
})
export class PageComponent {
  constructor(private router:Router) {

  }
}
