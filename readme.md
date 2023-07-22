project/
|-- cmd/              // ประกอบด้วยไฟล์ที่เกี่ยวข้องกับการ run แอปพลิเคชันหรือคำสั่งเริ่มต้นต่าง ๆ
|
|-- configs/          // เก็บไฟล์ตั้งค่าที่เกี่ยวข้องกับระบบของแอปพลิเคชัน เช่น การตั้งค่าฐานข้อมูล
|
|-- docs/             // เก็บเอกสารต่าง ๆ เช่น API Document ที่สร้างขึ้นโดย Swag Go
|   └── swagger/      // เก็บไฟล์ API Document ที่สร้างขึ้นโดย Swag Go
|
|-- internal/         // เก็บโค้ดที่สำคัญและเป็นส่วนหลักของแอปพลิเคชัน
|   ├── app/          // เก็บส่วนที่เกี่ยวข้องกับการทำธุรกิจของแอปพลิเคชัน
|   |   ├── entities/ // เก็บโครงสร้างข้อมูล (Data Model) ที่เกี่ยวข้องกับธุรกิจของแอปพลิเคชัน
|   |   |   ├── models/ // เก็บโครงสร้างข้อมูล (Data Model) ที่ใช้ในการเก็บข้อมูลในฐานข้อมูล
|   |   |   └── schema/ // เก็บโครงสร้างข้อมูล (Data Model) ที่ใช้ในการติดต่อกับผู้ใช้ผ่าน API
|   |   |
|   |   ├── repositories/ // เก็บโค้ดที่เกี่ยวข้องกับการเข้าถึงฐานข้อมูล (Data Access Layer)
|   |   └── usecases/  // เก็บโค้ดที่เกี่ยวข้องกับการจัดการธุรกิจหลักของแอปพลิเคชัน (Business Logic)
|   |
|   ├── handlers/     // เก็บส่วนที่เกี่ยวข้องกับการติดต่อกับผู้ใช้และตอบกลับผลลัพธ์
|   |   ├── middleware/ // เก็บโค้ด Middleware ที่ใช้ในการกระทำก่อนหรือหลังการทำงานของ Handlers
|   |   └── router/     // เก็บโค้ดการกำหนดเส้นทางการเรียกใช้ Handlers
|   |
|   └── infrastructures/ // เก็บโค้ดที่เกี่ยวข้องกับระบบในการรองรับแอปพลิเคชัน เช่น Database, Web Framework
|       ├── database/   // เก็บโค้ดการเชื่อมต่อกับฐานข้อมูล
|       └── web_framework/ // เก็บโค้ดที่เกี่ยวข้องกับการตั้งค่าและใช้งาน Web Framework
|
|-- readme.md          // ไฟล์ที่ใช้ในการเก็บคำอธิบายและคำแนะนำต่าง ๆ เกี่ยวกับโปรเจกต์