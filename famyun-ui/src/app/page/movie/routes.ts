import {Routes} from "@angular/router";
import {MovieListComponent} from "./movie-list/movie-list.component";
import {MovieEditComponent} from "./movie-edit/movie-edit.component";
import {MovieUploadComponent} from "./movie-upload/movie-upload.component";
import {MovieComponent} from "./movie.component";

export const routes:Routes = [
  {
    path:'',
    component:MovieComponent,
    children:[
      {path:'',redirectTo:'list',pathMatch:'full'},
      {path:'list',component:MovieListComponent},
      {path:'edit',component:MovieEditComponent},
      {path:'add',component:MovieUploadComponent},
    ]
  }
]
