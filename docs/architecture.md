## 技術構成

| 領域 | 技術 |
| --- | --- |
| Frontend | Next.js / TypeScript |
| Backend | Go |
| Database | MySQL |
| AI API | Python / FastAPI |
| Vector Database | PostgreSQL / pgvector |
| Infrastructure | AWS |
| IaC | Terraform |
| CI/CD | GitHub Actions |
| Local Environment | Docker / Docker Compose |

## システム構成

```mermaid
flowchart TD
    User[客 / 店舗スタッフ] --> Frontend[Next.js]

    Frontend --> Backend[Go API]
    Backend --> MySQL[(MySQL)]

    Backend --> AI[FastAPI]
    AI --> PostgreSQL[(PostgreSQL + pgvector)]
    AI --> LLM[LLM API]

