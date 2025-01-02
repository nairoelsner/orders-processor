import { Module } from '@nestjs/common';
import { OrdersModule } from './orders/orders.module';
import { RabbitmqModule } from './rabbitmq/rabbitmq.module';

@Module({
  imports: [OrdersModule, RabbitmqModule],
  controllers: [],
  providers: [],
})
export class AppModule {}
