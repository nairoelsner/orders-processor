import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  console.log('Starting application');
  console.log('RABBITMQ_URI --->', process.env.RABBITMQ_URI);

  await app.listen(process.env.PORT ?? 3000);
}
bootstrap();
