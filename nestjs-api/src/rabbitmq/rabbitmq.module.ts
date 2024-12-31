import { RabbitMQModule } from '@golevelup/nestjs-rabbitmq';
import { Module } from '@nestjs/common';

@Module({
    imports: [
        RabbitMQModule.forRoot(RabbitMQModule, {
          exchanges: [
            {
              name: 'amq.topic',
              type: 'topic',
            },
          ],
          uri: 'amqp://guest:guest@localhost:5672',
          connectionInitOptions: { wait: false },
        }),
      ],
      exports: [RabbitMQModule],
})
export class RabbitmqModule {}
