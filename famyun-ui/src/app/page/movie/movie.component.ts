import {Component, OnInit} from '@angular/core';
import {CommonModule} from "@angular/common";
import {MovieEditComponent} from "./movie-edit/movie-edit.component";
import {MovieUploadComponent} from "./movie-upload/movie-upload.component";
import {MovieListComponent} from "./movie-list/movie-list.component";
import {RouterModule, Routes} from "@angular/router";

@Component({
  selector: 'app-movie',
  standalone:true,
  imports:[
    RouterModule,
    CommonModule,
    MovieEditComponent,
    MovieUploadComponent,
    MovieListComponent,
  ],
  templateUrl: './movie.component.html',
  styleUrls: ['./movie.component.scss']
})
export class MovieComponent implements OnInit {

  constructor() {
  }

  ngOnInit(): void {
  }

}


