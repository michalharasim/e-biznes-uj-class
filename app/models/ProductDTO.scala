package models

import play.api.libs.json._

case class ProductDTO(name: String, price: Double)

object ProductDTO {
  implicit val format: OFormat[ProductDTO] = Json.format[ProductDTO]
}