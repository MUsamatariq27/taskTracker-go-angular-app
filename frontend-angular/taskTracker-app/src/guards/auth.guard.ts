import { CanActivateFn, Router } from '@angular/router';
import { ApiService } from '../services/api-service/api.service';
import { inject } from '@angular/core';


export const authGuard: CanActivateFn = () => {
  const apiService = inject(ApiService);
  const router = inject(Router);

  if (apiService.isLoggedIn()) {
    return true;
  } else {
    router.navigate(['/login']);
    return false;
  }
}
