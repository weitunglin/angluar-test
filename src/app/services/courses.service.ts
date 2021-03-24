import { Injectable } from "@angular/core";
import { HttpClient, HttpParams } from "@angular/common/http";
import { Observable } from "rxjs";
import { map } from 'rxjs/operators';
import { Course } from "../model/course";
import { Lesson } from "../model/lesson";

@Injectable()
export class CoursesService {
  constructor(private http: HttpClient) {
  }

  findCourseById(courseId: number): Observable<Course> {
    return this.http.get<Course>(`/api/courses/${courseId}`);
  }

  findAllCourses(): Observable<Course[]> {
    return this.http.get(`/api/courses`)
      .pipe(
        map((res: any) => res.payload)
      );
  }

  findAllCourseLessons(courseId: number): Observable<Lesson[]> {
    return this.http.get(`/api/lessons`, {
      params: new HttpParams()
        .set('courseId', courseId.toString())
        .set('pageNumber', '0')
        .set('pageSize', '100')
    }).pipe(
      map((res: any) => res.payload)
    );
  }

  findLessons(courseId: number, filter: string = '', sortOrder: string = 'asc', pageNumber: number = 0, pageSize: number = 3): Observable<Lesson[]> {
    return this.http.get(`/api/lessons`, {
      params: new HttpParams()
        .set('courseId', courseId.toString())
        .set('filter', filter)
        .set('sortOrder', sortOrder)
        .set('pageNumber', pageNumber.toString())
        .set('pageSize', pageSize.toString())
    }).pipe(
      map((res: any) => res.payload)
    );
  }
}
