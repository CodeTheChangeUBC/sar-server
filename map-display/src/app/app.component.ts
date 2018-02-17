import { Component } from '@angular/core';
import mapboxgl from 'mapbox-gl';

mapboxgl.accessToken = 'undefined';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})

export class AppComponent {
  title = 'BC Search and Rescue Field Server';
}
