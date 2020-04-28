import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { MessageService } from './message.service';
import { Item } from './item';

@Injectable({
  providedIn: 'root'
})
export class ItemService {

  private itemsUrl = '/items/';
  private itemUrl = '/item/';
  private itemUpdateUrl = '/item/update/';
  private itemCreateUrl = '/item/create/';
  private itemDeleteUrl = '/item/delete/';

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient,
    private messageService: MessageService) { }

  private log(message: string) {
    this.messageService.add(`HeroService: ${message}`);
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      console.error(error);
      this.log(`${operation} failed: ${error.message}`);
      return of(result as T);
    };
  }

  getItem(id: number): Observable<Item> {
    const url = `${this.itemUrl}?id=${id}`;
    return this.http.get<Item>(url)
      .pipe(
        tap(_ => this.log(`Fetched item id: ${id}`)),
        catchError(this.handleError<Item>(`getItem id=${id}`))
      );
  }

  getItems(): Observable<Item[]> {
    return this.http.get<Item[]>(this.itemsUrl)
      .pipe(
        tap(_ => this.log(`Fetched items`)),
        catchError(this.handleError<Item[]>('getItems', []))
      );
  }

  updateItem(item: Item): Observable<any> {
    return this.http.put(this.itemUpdateUrl, item, this.httpOptions)
      .pipe(
        tap(_ => this.log(`Updated item: ${item.Description}`)),
        catchError(this.handleError<any>('updateItem'))
      );
  }

  addItem(item: Item): Observable<Item> {
    return this.http.post<Item>(this.itemCreateUrl, item, this.httpOptions)
      .pipe(
        tap((newItem: Item) => this.log(`Added item with priority: ${newItem.Priority}`)),
        catchError(this.handleError<Item>('addItem'))
      );
  }

  deleteItem(item: Item | number): Observable<Item> {
    const id = typeof item === 'number' ? item : item.Id;
    const url = `${this.itemDeleteUrl}?id=${id}`;
  
    return this.http.delete<Item>(url, this.httpOptions)
      .pipe(
        tap(_ => this.log(`Deleted item with id: ${id}`)),
        catchError(this.handleError<Item>('deleteItem'))
      );
  }

}
