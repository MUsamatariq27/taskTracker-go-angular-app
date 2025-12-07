import { Component } from '@angular/core';
import { ApiService } from '../services/api-service/api.service';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';

@Component({
  selector: 'app-header',
  imports: [CommonModule],
  templateUrl: './header.component.html',
  styleUrl: './header.component.css'
})
export class HeaderComponent {

 constructor(public apiService: ApiService, private router: Router) {}

  logout() {
    this.apiService.logout();
    this.router.navigate(['/login'])
  }

}
