import { Injectable } from '@angular/core';
import { environment } from '../environments/environment';

@Injectable()
export class LoggerService {

  constructor() { }

  trace(message: string): void {
    if (!environment.production) {
      console.log(`[TRACE] ${message}`);
    }
  }

  debug(message: string): void {
    if (!environment.production) {
      console.log(`[DEBUG] ${message}`);
    }
  }

  info(message: string): void {
    console.log(`[INFO] ${message}`);
  }

  error(message: string): void {
    console.log(`[ERROR] ${message}`);
  }
}
